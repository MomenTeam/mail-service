package models

import (
	"context"

	"github.com/MomenTeam/consumer-ms/database"
	"gopkg.in/mgo.v2/bson"
)

type MailTemplate struct {
	Template string `bson:"template"`
}

func ReadTemplate(mailType int) (string, error) {
	var mailTemplate MailTemplate
	err := database.MailTemplatesCollection.FindOne(context.TODO(), bson.M{"mailTemplateType": mailType}).Decode(&mailTemplate)
	return mailTemplate.Template, err
}
