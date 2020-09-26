package main

import (
	"github.com/MomenTeam/consumer-ms/database"
	"github.com/MomenTeam/consumer-ms/routes"
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
