package handler

import (
	"fmt"
	"net/http"
	"start-up-rh/auth"
	"start-up-rh/helper"
	"start-up-rh/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.ErrorValidationMessege(err)
		errormessege := gin.H{"errors": errors}

		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errormessege)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newuser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GeneredToken(newuser.Id)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newuser, token)

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {

	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationMessege(err)
		errormessege := gin.H{"errors": errors}

		response := helper.ApiResponse("Login account failed", http.StatusUnprocessableEntity, "error", errormessege)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	logginUsers, err := h.userService.Login(input)
	if err != nil {
		errormessege := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login account failed", http.StatusUnprocessableEntity, "error", errormessege)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GeneredToken(logginUsers.Id)
	if err != nil {
		response := helper.ApiResponse("Login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(logginUsers, token)

	response := helper.ApiResponse("Succestfully login", http.StatusOK, "succest", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) CheckEmail(c *gin.Context) {
	//user input login
	//email yg di inut di teruskan ke strack input
	// stack input mempassing ke service
	// dari service akan memanggil repository  untuk pengecekan email
	//reposytory menghubungkan ke db (database)
	var CheckEmail user.CheckEmailInput
	err := c.ShouldBindJSON(&CheckEmail)
	if err != nil {
		errors := helper.ErrorValidationMessege(err)
		errormessege := gin.H{"errors": errors}

		response := helper.ApiResponse("E-mail checking failed", http.StatusUnprocessableEntity, "error", errormessege)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	emailAvailable, err := h.userService.CheckEmailAvailable(CheckEmail)

	if err != nil {
		errormessege := gin.H{"errors": "Server is error"}

		response := helper.ApiResponse("E-mail checking failed", http.StatusUnprocessableEntity, "error", errormessege)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{
		"is_avalible": emailAvailable,
	}

	var metaMessages string

	if emailAvailable {
		metaMessages = "Email is available"
	} else {
		metaMessages = "Email has bee registered"
	}

	response := helper.ApiResponse(metaMessages, http.StatusOK, "succest", data)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatas(c *gin.Context) {

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Faile to upload avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
	}

	CurrentUser := c.MustGet("currentUser").(user.User)
	UserId := CurrentUser.Id

	path := fmt.Sprintf("images/%d-%s", UserId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Faile to upload avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
	}

	_, err = h.userService.SaveAvatar(UserId, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Faile to upload avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Succest uploaded avatar", http.StatusOK, "succest", data)

	c.JSON(http.StatusOK, response)

}
