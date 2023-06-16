package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaign) (Campaign, error)
	UpdateCampaign(inputId GetCampaignDetailInput, inputData CreateCampaign) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {

	if userId != 0 {
		campaigns, err := s.repository.FindById(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindIdCam(input.Id)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaign) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserId = input.User.Id

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.Id)
	campaign.Slug = slug.Make(slugCandidate)

	NewCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return NewCampaign, err
	}
	return NewCampaign, nil
}

func (s *service) UpdateCampaign(inputId GetCampaignDetailInput, inputData CreateCampaign) (Campaign, error) {
	campaign, err := s.repository.FindIdCam(inputId.Id)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != inputData.User.Id {
		return campaign, errors.New("Not on owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	UpdateCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return UpdateCampaign, err
	}
	return UpdateCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindIdCam(input.CampaignId)
	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserId != input.User.Id {
		return CampaignImage{}, errors.New("Not on owner of the campaign")
	}

	isPrimary := 0

	if input.IsPrimary {
		isPrimary = 1

		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignId)

		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignId = input.CampaignId
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	NewCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return NewCampaignImage, err
	}
	return NewCampaignImage, nil
}
