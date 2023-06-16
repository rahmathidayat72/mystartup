package handler

import (
	"fmt"
	"net/http"
	"start-up-rh/campaign"
	"start-up-rh/helper"
	"start-up-rh/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CampaignHandler struct {
	sevice campaign.Service
}

func NewCampaingHandler(service campaign.Service) *CampaignHandler {
	return &CampaignHandler{service}
}

//api/vi/campaigns

func (h *CampaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.sevice.GetCampaigns(userId)
	if err != nil {
		response := helper.ApiResponse("Error to get campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List to campaigns", http.StatusOK, "succest", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *CampaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.sevice.GetCampaignById(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Campaign Detai", http.StatusOK, "succesc", campaign.FormatDetailCampaign(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaign
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.ErrorValidationMessege(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to create campaign", http.StatusBadRequest, "error", errorMassage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := h.sevice.CreateCampaign(input)
	if err != nil {
		response := helper.ApiResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Create campaign", http.StatusOK, "Succect", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)

}

func (h *CampaignHandler) UpdateCampaign(c *gin.Context) {
	var inputId campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.ApiResponse("Failed to update  campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData campaign.CreateCampaign

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.ErrorValidationMessege(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to update campaign", http.StatusBadRequest, "error", errorMassage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	inputData.User = currentUser

	updateCampaign, err := h.sevice.UpdateCampaign(inputId, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to update  campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Succect to update campaign", http.StatusOK, "Succect", campaign.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *CampaignHandler) UploadImage(c *gin.Context) {

	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.ErrorValidationMessege(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.ApiResponse("Failed to update campaign image", http.StatusBadRequest, "error", errorMassage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Faile to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
	}

	CurrentUser := c.MustGet("currentUser").(user.User)
	UserId := CurrentUser.Id

	path := fmt.Sprintf("images/%d-%s", UserId, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Faile to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
	}
	_, err = h.sevice.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Faile to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Succest uploaded campaign image", http.StatusOK, "succest", data)

	c.JSON(http.StatusOK, response)

}
