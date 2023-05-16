package handler

import (
	"net/http"
	"start-up-rh/helper"
	"start-up-rh/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
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
	formatter := user.FormatUser(newuser, "initoken")

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

	LogginUsers, err := h.userService.Login(input)
	if err != nil {
		errormessege := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login account failed", http.StatusUnprocessableEntity, "error", errormessege)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(LogginUsers, "token")

	response := helper.ApiResponse("Succestfully loggin", http.StatusOK, "succest", formatter)

	c.JSON(http.StatusOK, response)

}
