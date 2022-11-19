package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
)

func SendSimpleMessage(c *gin.Context) (string, error) {
	mg := mailgun.NewMailgun(os.Getenv("mailgun_domain"), os.Getenv("mailgun_apikey"))
	mg.SetAPIBase(os.Getenv("mailgun_apibase"))
	m := mg.NewMessage(
		"John <"+os.Getenv("mailgun_sender")+">",
		"Hello",
		"Testing some Mailgun awesomeness!",
		"receiver@example.com",
	)
	_, id, err := mg.Send(c, m)
	return id, err
}
