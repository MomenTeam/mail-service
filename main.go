package main

import (
	"github.com/MomenTeam/mail-service/database"
	"github.com/MomenTeam/mail-service/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	database.Setup()
}

func main() {
	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
