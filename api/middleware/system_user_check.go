package middleware

import (
	"api/database"
	"api/models"
	"api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsSystemUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get("authScope")
		authScope := data.(structs.AuthScope)

		var user models.User
		database.DB.First(&user, authScope.UserID)

		if !user.IsSystemUser {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "You do not have the appropriate permissions to access this data.", "data": nil})
			return
		}

		c.Next()
	}
}
