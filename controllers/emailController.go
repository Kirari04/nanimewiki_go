package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
)

func SendMessage(c *gin.Context, header string, body string, receiver string) (string, error) {
	mg := mailgun.NewMailgun(os.Getenv("mailgun_domain"), os.Getenv("mailgun_apikey"))
	mg.SetAPIBase(os.Getenv("mailgun_apibase"))
	m := mg.NewMessage(
		os.Getenv("mailgun_sender_name")+" <"+os.Getenv("mailgun_sender")+">",
		header,
		"",
		receiver,
	)
	m.SetHtml(body)
	_, id, err := mg.Send(c, m)
	return id, err
}
