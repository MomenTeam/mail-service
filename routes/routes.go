package routes

import (
	"net/http"

	"github.com/MomenTeam/mail-service/controllers"
	"github.com/gin-gonic/gin"
)

// Routes func
func Routes(router *gin.Engine) {
	router.GET("/ping", ping)
	router.NoRoute(notFound)

	email := router.Group("/v1/email")
	{
		email.POST("/", controllers.SendEmail)

	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "pong",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
