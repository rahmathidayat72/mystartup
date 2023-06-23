package transaction

import (
	"errors"
	"start-up-rh/campaign"
)

type Service interface {
	GetTransactionByCampaignId(input GetTransactionCampaignInput) ([]Transactions, error)
	GetTransactionByUserId(userId int) ([]Transactions, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignId(input GetTransactionCampaignInput) ([]Transactions, error) {

	campaign, err := s.campaignRepository.FindIdCam(input.Id)
	if err != nil {
		return []Transactions{}, err
	}
	if campaign.UserId != input.User.Id {
		return []Transactions{}, errors.New("Not an owner of the campaign")
	}

	transaction, err := s.repository.GetByCampaignId(input.Id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) GetTransactionByUserId(userId int) ([]Transactions, error) {

	transaction, err := s.repository.GetByUserId(userId)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
