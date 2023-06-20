package transaction

import "start-up-rh/user"

type GetTransactionCampaignInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}
