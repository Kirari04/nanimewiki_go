package controllers

import (
	"ch/kirari/animeApi/models"
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"success": 0,
			"error":   "Bad Request",
			"data":    nil,
			"len":     nil,
		})
		return
	}
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": 0,
				"error":   "The field " + err.Field() + " is invalid.",
				"data":    nil,
				"len":     nil,
			})
			return
		}
	}

	//check username exists
	var countUsername int64
	res_username := models.DB.Where(&models.User{Username: req.Username}).Model(&models.User{}).Count(&countUsername)
	if res_username.Error != nil {
		InternalServerError_response(c)
		return
	}
	if countUsername > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"success": 0,
			"error":   "This Username is already taken.",
			"data":    nil,
			"len":     nil,
		})
		return
	}

	//check email exists
	var countEmail int64
	res_email := models.DB.Where(&models.User{Email: req.Email}).Model(&models.User{}).Count(&countEmail)
	if res_email.Error != nil {
		InternalServerError_response(c)
		return
	}
	if countEmail > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"success": 0,
			"error":   "This Email is already taken.",
			"data":    nil,
			"len":     nil,
		})
		return
	}
	// validation finished

	passwordHash, err := HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": 0,
			"error":   "The server can't process this password. Please choose another one.",
			"data":    nil,
			"len":     nil,
		})
		return
	}

	res_user := models.DB.Create(&models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: passwordHash,
	})

	if res_user.Error != nil {
		InternalServerError_response(c)
		return
	}
	header := ""
	message := ""
	id, err_mail := SendMessage(c, header, message, req.Email)
	if err_mail != nil || id == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": 0,
			"error":   "Couldn't send verification e-mail. Please contact support.",
			"data":    nil,
			"len":     nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": 1,
		"error":   nil,
		"data":    nil,
		"len":     nil,
	})
}
