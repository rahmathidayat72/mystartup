package transaction

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignId(campaignID int) ([]Transactions, error)
	GetByUserId(userId int) ([]Transactions, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignId(campaignID int) ([]Transactions, error) {
	var transaction []Transactions

	err := r.db.Preload("User").Where("campaign_id= ?", campaignID).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) GetByUserId(userId int) ([]Transactions, error) {
	var transaction []Transactions

	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary=1").Where("user_id= ?", userId).Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
