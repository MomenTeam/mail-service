package main

import (
	"github.com/MomenTeam/consumer-ms/routes"
	"github.com/gin-gonic/gin"
)

const (
	QueueURL    = "https://sqs.eu-central-1.amazonaws.com/070835381129/mail-queue"
	Region      = "eu-central-1"
	CredPath    = "./.credentials"
	CredProfile = "aws-cred-profile"
)

func main() {

	r := gin.Default()
	routes.Routes(r)
	r.Run()
}
