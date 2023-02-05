package controllers

import (
	"api/config"
	"api/initializers/database"
	"api/models"
	"api/services/auth_service"
	"api/services/users_service"
	"fmt"
	"log"

	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	type LoginInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var input LoginInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on login request.", "data": err})
		return
	}

	var user models.User

	database.DB.Where(&models.User{Email: input.Email}).First(&user)

	if user.ID.String() != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "User not found.", "data": nil})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Incorrect email and password combination."})
		return
	}

	tokens, err := auth_service.CreateTokens(user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error on register request.", "data": err})
		return
	}

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error on login request.", "data": err})
		return
	}

	c.SetCookie("ucid", tokens.RefreshToken, int((time.Now().Add(time.Hour * 24 * 14)).Unix()), "/api/refresh_token", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User logged in.", "data": tokens.AccessToken})
}

func Register(c *gin.Context) {
	type RegisterInput struct {
		FirstName            string `json:"firstName" binding:"required"`
		LastName             string `json:"lastName" binding:"required"`
		Email                string `json:"email" binding:"required"`
		Password             string `json:"password" binding:"required"`
		PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
	}

	var input RegisterInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on register request.", "data": err})
		return
	}

	if input.Password != input.PasswordConfirmation {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Password and Confirmation do not match.", "data": nil})
		return
	}

	var existingUser models.User

	database.DB.Where(&models.User{Email: input.Email}).Find(&existingUser)

	if existingUser.ID.String() != "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error on register request.", "data": nil})
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error on register request.", "data": err})
		return
	}

	c.SetCookie("ucid", tokens.RefreshToken, int((time.Now().Add(time.Hour * 24 * 14)).Unix()), "/api/refresh_token", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User registered", "data": tokens.AccessToken})
}

func RefreshToken(c *gin.Context) {
	token, err := c.Cookie("ucid")

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"token": ""})
		return
	}

	verified_token, jwt_err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Get("REFRESH_SECRET")), nil
	})

	if jwt_err != nil {
		c.JSON(http.StatusOK, gin.H{"token": ""})
		return
	}

	claims := verified_token.Claims.(jwt.MapClaims)

	user_id, _ := uuid.Parse(claims["iss"].(string))

	user, get_user_err := users_service.GetUserById(user_id)

	if get_user_err != nil {
		c.JSON(http.StatusOK, gin.H{"token": ""})
		return
	}

	tokens, tokens_err := auth_service.CreateTokens(*user)

	if tokens_err != nil {
		c.JSON(http.StatusOK, gin.H{"token": ""})
		return
	}

	c.SetCookie("ucid", tokens.RefreshToken, int((time.Now().Add(time.Hour * 24 * 14)).Unix()), "/api/refresh_token", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"token": tokens.AccessToken})
}
