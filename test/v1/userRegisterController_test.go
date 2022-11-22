package controllers__test

import (
	"ch/kirari/animeApi/models"
	"ch/kirari/animeApi/setups"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

type tmpEmailListResponse struct {
	ID      int    `json:"id"`
	From    string `json:"from"`
	Subject string `json:"subject"`
}

func TestUserRegisterCorrect(t *testing.T) {
	username := RandStringBytesRmndr(10)
	password := RandStringBytesRmndr(10)
	t.Cleanup(func() {
		setups.ConnectDatabase(os.Getenv("database"))
		setups.DB.Unscoped().Where(&models.User{Username: username}).Delete(&models.User{})
	})
	// generate temporary email
	var emailList []string
	getTmpEmails(t, os.Getenv("tmpmail_1secmail_getmail"), 200, &emailList)
	AssertEqualInt(t, len(emailList), 1)

	// sign up
	res := postReq(t, "/api/v1/user/register", 200, url.Values{
		"email":    {emailList[0]},
		"username": {username},
		"password": {password},
	})
	AssertEqualInt(t, res.Success, 1)

	// check if message retreived
	emailData := strings.Split(emailList[0], "@")
	uri := os.Getenv("tmpmail_1secmail_getmsg") + "login=" + emailData[0] + "&domain=" + emailData[1]
	trysLeft := 60 // wait up to two minute
	sleepSeconds := time.Second * 2
	messageFound := false
	t.Logf("Checking on url: %v", uri)
	for i := 0; i < trysLeft; i++ {
		var messageList []tmpEmailListResponse
		getTmpMessages(t, uri, 200, &messageList)
		for _, v := range messageList {
			if strings.Contains(v.Subject, username) {
				messageFound = true
				break
			}
		}
		if messageFound {
			break
		}
		t.Logf("No match in check %v", i)
		time.Sleep(sleepSeconds)
	}
	AssertEqualBool(t, messageFound, true)
}
