package controllers

import (
	"ch/kirari/animeApi/models"
	"ch/kirari/animeApi/setups"
	"ch/kirari/animeApi/templates"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MapRequest_UserRegister struct {
	Email    string `form:"email" validate:"required,email"`
	Username string `form:"username" validate:"required,min=2,max=22"`
	Password string `form:"password" validate:"required,min=8,max=128"`
}

func UserRegister(c *gin.Context) {
	var req MapRequest_UserRegister
	if c.ShouldBind(&req) != nil {
		CustomError_response(c, http.StatusBadRequest, "Bad Request")
		return
	}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			CustomError_response(c, http.StatusBadRequest, "The field "+err.Field()+" is invalid.")
			return
		}
	}

	//check username exists
	var countUsername int64
	res_username := setups.DB.Where(&models.User{Username: req.Username}).Model(&models.User{}).Count(&countUsername)
	if res_username.Error != nil {
		InternalServerError_response(c)
		return
	}
	if countUsername > 0 {
		CustomError_response(c, http.StatusConflict, "This Username is already taken.")
		return
	}

	//check email exists
	var countEmail int64
	res_email := setups.DB.Where(&models.User{Email: req.Email}).Model(&models.User{}).Count(&countEmail)
	if res_email.Error != nil {
		InternalServerError_response(c)
		return
	}
	if countEmail > 0 {
		CustomError_response(c, http.StatusConflict, "This Email is already taken.")
		return
	}

	// hash password
	passwordHash, err := HashPassword(req.Password)
	if err != nil {
		CustomError_response(c, http.StatusBadRequest, "The server can't process this password. Please choose another one.")
		return
	}

	// add user
	res_user := setups.DB.Create(&models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: passwordHash,
	})

	if res_user.Error != nil {
		InternalServerError_response(c)
		return
	}

	// load email template
	emailTemplate, err := templates.Get(templates.Data.EmailRegister, "de")
	if err != nil {
		InternalServerError_response(c)
		return
	}

	// template params
	vars := []models.TemplateVars{
		{
			Variable: "username",
			Value:    req.Username,
		},
		{
			Variable: "service_name",
			Value:    os.Getenv("service_name"),
		},
		{
			Variable: "service_domain",
			Value:    os.Getenv("service_domain"),
		},
		{
			Variable: "code",
			Value:    "187Test",
		},
	}

	// build email
	header := templates.Prepare(emailTemplate["head"], vars)
	message := templates.Prepare(emailTemplate["body"], vars)

	// send email
	id, err_mail := SendMessage(c, header, message, req.Email)
	if err_mail != nil || id == "" {
		log.Printf("Error message: %v / Email id %v\n", err_mail.Error(), id)
		CustomError_response(c, http.StatusInternalServerError, "Couldn't send verification e-mail. Please contact support.")
		return
	}

	OK_response(c, nil, nil)
}
