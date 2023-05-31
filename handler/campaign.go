package handler

import (
	"net/http"
	"start-up-rh/campaign"
	"start-up-rh/helper"
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

	campaign, err := h.sevice.GetCampaignById(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Campaign Detai", http.StatusOK, "succesc", campaign)
	c.JSON(http.StatusOK, response)
}
