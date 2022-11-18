package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
)

func SendSimpleMessage(c *gin.Context, domain, apiKey string) (string, error) {
	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		"Excited User <mailgun@YOUR_DOMAIN_NAME>",
		"Hello",
		"Testing some Mailgun awesomeness!",
		"YOU@YOUR_DOMAIN_NAME",
	)
	_, id, err := mg.Send(c, m)
	return id, err
}
