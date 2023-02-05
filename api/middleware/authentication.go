package middleware

import (
	"api/config"
	"api/structs"

	"fmt"
	"net/http"
  "github.com/google/uuid"
  "github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const BEARER_PREFIX = "Bearer "

func Protected() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "No authorization header passed.", "data": nil})
			return
		}

		tokenString := authHeader[len(BEARER_PREFIX):]

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.Get("ACCESS_SECRET")), nil
		})

		claims := token.Claims.(jwt.MapClaims)

		userId, _ := claims["iss"].(string)

    parsedUserID, uuid_err := uuid.Parse(userId)

    if uuid_err != nil {
      c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "Invalid userId", "data": nil})
    }

		if token.Valid {
			authScope := structs.AuthScope{
				UserID: parsedUserID,
			}
			c.Set("authScope", authScope)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Bad token.", "data": nil})
			return
		}
	}
}
