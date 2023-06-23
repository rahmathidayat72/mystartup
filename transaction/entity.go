package transaction

import (
	"start-up-rh/campaign"
	"start-up-rh/user"
	"time"
)

type Transactions struct {
	Id         int
	CampaignId int
	UserId     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
