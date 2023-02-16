package controllers

import (
	"api/config"
	"api/initializers/database"
	"api/models"
	"api/services/auth_service"
	"api/services/users_service"
  "api/structs"
	"fmt"
	"log"

	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
  Email    string `json:"email" binding:"required"`
  Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
  FirstName            string `json:"firstName" binding:"required"`
  LastName             string `json:"lastName" binding:"required"`
  Email                string `json:"email" binding:"required"`
  Password             string `json:"password" binding:"required"`
  PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error on login request."})
		return
	}

	var user models.User

	database.DB.Where(&models.User{Email: input.Email}).First(&user)

	if user.ID != uuid.Nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "User not found."})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Incorrect email and password combination."})
		return
	}

	tokens, err := auth_service.CreateTokens(user)

	if err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Error on register request."})
		return
	}

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Error on login request."})
		return
	}

	c.SetCookie("ucid", tokens.RefreshToken, int((time.Now().Add(time.Hour * 24 * 14)).Unix()), "/api/refresh_token", "localhost", false, true)

	c.JSON(http.StatusOK, tokens.AccessToken)
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.BindJSON(&input); err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error on register request."})
		return
	}

	if input.Password != input.PasswordConfirmation {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Password and Confirmation do not match."})
		return
	}

	var existingUser models.User

	database.DB.Where(&models.User{Email: input.Email}).Find(&existingUser) 

	if existingUser.ID != uuid.Nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Error on register request."})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  string(password),
	}

	database.DB.Create(&user)

	tokens, err := auth_service.CreateTokens(user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Error on register request."})
		return
	}

	c.SetCookie("ucid", tokens.RefreshToken, int((time.Now().Add(time.Hour * 24 * 14)).Unix()), "/api/refresh_token", "localhost", false, true)

	c.JSON(http.StatusOK, tokens.AccessToken)
}

func RefreshToken(c *gin.Context) {
	token, err := c.Cookie("ucid")

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, "")
		return
	}

	verified_token, jwt_err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Get("REFRESH_SECRET")), nil
	})

	if jwt_err != nil {
		c.JSON(http.StatusOK, "")
		return
	}

	claims := verified_token.Claims.(jwt.MapClaims)

	user_id, _ := uuid.Parse(claims["iss"].(string))

	user, get_user_err := users_service.GetUserById(user_id)

	if get_user_err != nil {
		c.JSON(http.StatusOK, "")
		return
	}

	tokens, tokens_err := auth_service.CreateTokens(*user)

	if tokens_err != nil {
		c.JSON(http.StatusOK, "")
		return
	}

	c.SetCookie("ucid", tokens.RefreshToken, int((time.Now().Add(time.Hour * 24 * 14)).Unix()), "/api/refresh_token", "localhost", false, true)

	c.JSON(http.StatusOK, tokens.AccessToken)
}
