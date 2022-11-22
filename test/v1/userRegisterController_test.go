package controllers__test

import (
	"ch/kirari/animeApi/helpers"
	"ch/kirari/animeApi/models"
	"ch/kirari/animeApi/setups"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

type tmpEmailListResponse struct {
	ID      int    `json:"id"`
	From    string `json:"from"`
	Subject string `json:"subject"`
}

type tmpEmailResponse struct {
	ID       int    `json:"id"`
	From     string `json:"from"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
	TextBody string `json:"textBody"`
	HtmlBody string `json:"htmlBody"`
}

func TestUserRegisterCorrect(t *testing.T) {
	username := helpers.RandStringBytesRmndr(4)
	password := helpers.RandStringBytesRmndr(10)
	// cleens user up after test
	t.Cleanup(func() {
		setups.ConnectDatabase(os.Getenv("database"))
		setups.DB.Unscoped().Where("1 = 1").Delete(&models.EmailVerificationCode{})
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
	loginUri := "login=" + emailData[0] + "&domain=" + emailData[1]
	uri := os.Getenv("tmpmail_1secmail_getmsgs") + loginUri
	trysLeft := 60 // wait up to two minute
	sleepSeconds := time.Second * 2
	messageFound := false
	var messageId int
	t.Logf("Checking on url: %v", uri)
	for i := 0; i < trysLeft; i++ {
		var messageList []tmpEmailListResponse
		getTmpMessages(t, uri, 200, &messageList)
		for _, v := range messageList {
			if strings.Contains(v.Subject, username) {
				messageId = v.ID
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

	// gert verification code from db
	setups.ConnectDatabase(os.Getenv("database"))
	var userData models.User
	if err := setups.DB.Where(&models.User{Username: username}).First(&userData).Error; err != nil {
		t.Fatal(err)
	}
	log.Println("UserID", userData.ID)

	var userKeyData models.EmailVerificationCode
	if err := setups.DB.Debug().Where(&models.EmailVerificationCode{UserID: userData.ID}).First(&userKeyData).Error; err != nil {
		t.Fatal(err)
	}

	verificationCode := userKeyData.Code
	log.Println("verificationCode", verificationCode)

	//check if message contains correct code
	messageUri := os.Getenv("tmpmail_1secmail_getmsg") + loginUri + "&id=" + strconv.Itoa(messageId)
	var tmpMessage tmpEmailResponse
	getTmpMessage(t, messageUri, 200, &tmpMessage)
	if !strings.Contains(tmpMessage.TextBody, "Code: "+verificationCode) {
		t.Fatalf("Verification Code not found inside text message:\n %v\n%v\n", tmpMessage.TextBody, verificationCode)
	}
	if !strings.Contains(tmpMessage.HtmlBody, ">"+verificationCode+"</code>") {
		t.Fatalf("Verification Code not found inside html message:\n %v\n%v\n", tmpMessage.TextBody, verificationCode)
	}
}
