package handler

import (
	"net/http"
	"start-up-rh/helper"
	"start-up-rh/transaction"
	"start-up-rh/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	Service transaction.Service
}

func NewTransactionHandler(Service transaction.Service) *transactionHandler {
	return &transactionHandler{Service}
}

func (h *transactionHandler) GetTransactionCampaign(c *gin.Context) {
	var input transaction.GetTransactionCampaignInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign's trasactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := h.Service.GetTransactionByCampaignId(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign's trasactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign's transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)

}

func (h *transactionHandler) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.Id

	transactions, err := h.Service.GetTransactionByUserId(userId)
	if err != nil {
		response := helper.ApiResponse("Failed to get user trasactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("User transaction", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)

}
