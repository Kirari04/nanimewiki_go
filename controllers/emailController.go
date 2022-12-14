package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
)

func SendMessage(c *gin.Context, header string, bodyTxt string, bodyHtml string, receiver string) (string, error) {
	mg := mailgun.NewMailgun(os.Getenv("mailgun_domain"), os.Getenv("mailgun_apikey"))
	mg.SetAPIBase(os.Getenv("mailgun_apibase"))
	m := mg.NewMessage(
		os.Getenv("mailgun_sender_name")+" <"+os.Getenv("mailgun_sender")+">",
		header,
		bodyTxt,
		receiver,
	)
	m.SetHtml(bodyHtml)
	_, id, err := mg.Send(c, m)
	return id, err
}
