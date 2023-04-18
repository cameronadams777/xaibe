package controllers

import (
	"api/config"
	"api/models"
	"api/services/stripe_service"
	"api/services/users_service"
	"api/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateNewStripeCustomerInput struct {
  BillingEmail    string   `binding:"required" json:"billingEmail"`
  BusinessName    string   `binding:"required" json:"businessName"`
  AddressLineOne  string   `binding:"required" json:"addressLineOne"`
  AddressLineTwo  string   `binding:"-" json:"addressLineTwo"`
  City            string   `binding:"required" json:"city"`
  State           string   `binding:"-" json:"state"`
  PostalCode      string   `binding:"-" json:"postalCode"`
  Country         string   `binding:"required" json:"country"`
}

type CreateNewTeamSubscriptionInput struct {
  NumberOfSeats   uint     `binding:"required" json:"numberOfSeats"`
}

type CreateNewTeamSubscriptionResponse struct {
  ClientSecret    string    `json:"clientSecret"`
}

func CreateNewStripeCustomer(c *gin.Context) {
  var input CreateNewStripeCustomerInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	current_user, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An unknown error occurred."})
		return
	}
 
  if len(current_user.StripeId) != 0 {
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "User already has stripe account."})
    return
  }

  stripe_customer, customer_create_err := stripe_service.CreateCustomer(stripe_service.CustomerMetadata{
    BillingEmail: input.BillingEmail,
    BusinessName: input.BusinessName,
    AddressLineOne: input.AddressLineOne,
    AddressLineTwo: input.AddressLineTwo,
    City: input.AddressLineTwo,
    State: input.State,
    PostalCode: input.PostalCode,
    Country: input.Country,
  })

  if customer_create_err != nil {
    fmt.Println(customer_create_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An unknown error occurred creating customer."})
    return
  }

  _, update_err := users_service.UpdateUser(authScope.UserID, models.User{ StripeId: stripe_customer.ID })

  if update_err != nil {
    fmt.Println(update_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An unknown error occurred creating customer."})
    return
  }

  c.JSON(http.StatusCreated, "")
}

func CreateNewTeamSubscription(c *gin.Context) {
   var input CreateNewTeamSubscriptionInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

  data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	current_user, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An unknown error occurred."})
		return
	}
 
  if len(current_user.StripeId) == 0 {
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "User does not have a account."})
    return
  }

  subscription, subscription_create_err := stripe_service.CreateSubscription(stripe_service.SubscriptionData{
    CustomerId: current_user.StripeId,
    PriceId: config.Get("STRIPE_TEAM_PRODUCT_ID"),
    Quantity: input.NumberOfSeats,
  })

  if subscription_create_err != nil {
    fmt.Println(subscription_create_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An unknown error occurred creating subscription."})
    return
  }

  c.JSON(http.StatusCreated, CreateNewTeamSubscriptionResponse{ClientSecret: subscription.LatestInvoice.PaymentIntent.ClientSecret})
}

