package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WebHook(c *gin.Context) {
	application := c.Param("application")

	if len(application) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Application not specified.", "data": nil})
		return
	}

	_, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error occurred reading in alert.", "data": err})
		return
	}

	// Take body data and push to redis under cache key
}
