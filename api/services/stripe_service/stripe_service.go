package stripe_service

import (
	"api/initializers/stripe_client"

	"github.com/stripe/stripe-go/v74"
)

type CustomerMetadata struct {
  BillingEmail string `binding:"required"`
  BusinessName string `binding:"required"`
  AddressLineOne string `binding:"required"`
  AddressLineTwo string `binding:"-"`
  City string `binding:"required"`
  State string `binding:"-"`
  PostalCode string `binding:"-"`
  Country string `binding:"required"`
}

type SubscriptionData struct {
  CustomerId    string
  PriceId       string
  Quantity      uint
  Metadata      map[string]string
}

func CreateCustomer(metadata CustomerMetadata) (*stripe.Customer, error) {
  params := &stripe.CustomerParams{
    Email: stripe.String(metadata.BillingEmail),
    Address: &stripe.AddressParams {
      Line1: stripe.String(metadata.AddressLineOne),
      Line2: stripe.String(metadata.AddressLineTwo),
      City: stripe.String(metadata.City),
      PostalCode: stripe.String(metadata.PostalCode),
      State: stripe.String(metadata.State),
      Country: stripe.String(metadata.Country),
    }, 
  }
  customer, err := stripe_client.StripeClient.Customers.New(params);

  if err != nil {
    return nil, err
  }

  return customer, nil
}

func CreateSubscription(metadata SubscriptionData) (*stripe.Subscription, error) {
  payment_settings := &stripe.SubscriptionPaymentSettingsParams{
    SaveDefaultPaymentMethod: stripe.String("on_subscription"),
  } 

  subscription_params := &stripe.SubscriptionParams{
    Customer: stripe.String(metadata.CustomerId),
    Items: []*stripe.SubscriptionItemsParams{
      {
        Price: stripe.String(metadata.PriceId),
        Quantity: stripe.Int64(int64(metadata.Quantity)),
      },
    },
    PaymentSettings: payment_settings,
    PaymentBehavior: stripe.String("default_incomplete"),
  }
  subscription_params.AddExpand("latest_invoice.payment_intent")

  for k, v := range metadata.Metadata {
    subscription_params.AddMetadata(k, v)
  }

  subscription, err := stripe_client.StripeClient.Subscriptions.New(subscription_params)

  if err != nil {
    return nil, err
  }

  return subscription, nil
}

func ConfirmCardPayment(payment_intent string, card_token string) error {
  params := &stripe.PaymentIntentConfirmParams{
    PaymentMethod: stripe.String(card_token),
  }
  _, err := stripe_client.StripeClient.PaymentIntents.Confirm(payment_intent, params)
  if err != nil {
    return err
  }
  return nil
}

// TODO: Fill this in
func UpdateSubscription() {

}

func CancelSubscription(subscription_id string) error {
  _, err := stripe_client.StripeClient.Subscriptions.Cancel(subscription_id, nil)
  if err != nil {
    return err
  }
  return nil
}
