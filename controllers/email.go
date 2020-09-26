package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/MomenTeam/consumer-ms/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

var (
	Host     = os.Getenv("EMAIL_HOST")
	Username = os.Getenv("EMAIL_USERNAME")
	Password = os.Getenv("EMAIL_PASSWORD")
)

// Information struct for mapping
type Information struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Type    int    `json:"type"`
}

func getEmailTemplate(mailType int) (string, error) {
	template, err := models.ReadTemplate(mailType)
	if err != nil {
		return "", err
	}
	return template, err
}

func send(information Information) (result Information, err error) {
	template, _ := models.ReadTemplate(information.Type)
	m := gomail.NewMessage()
	m.SetHeader("From", Username)
	m.SetHeader("To", information.Email)
	m.SetHeader("Subject", "Team Momentum.")
	m.SetBody("text/html", template)
	d := gomail.NewDialer(Host, 465, Username, Password)

	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Error while sending email %s", err)
		return information, err
	}
	return information, nil
}

// SendEmail function
func SendEmail(c *gin.Context) {
	information := &Information{}
	c.BindJSON(&information)

	result, err := send(*information)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Mail can't sent!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Mail successfully sent!",
		"data":    result,
	})
	return
}
