package middleware

import (
	"api/database"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("authScope")
		authScope := data.(models.AuthScope)

		var user models.User
		database.DB.First(&user, authScope.UserID)

		if !user.IsAdmin {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "You must have an admin role to interact with this route.", "data": nil})
			return
		}

		c.Next()
	}
}
