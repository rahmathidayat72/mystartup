package campaign

import (
	"strings"
)

type CampaignFormatter struct {
	Id              int    `json:"id"`
	UserId          int    `json:"user_id"`
	Name            string `json:"name"`
	ShotDescription string `json:"short_description"`
	Descrtiption    string `json:"description"`
	ImagesURL       string `json:"images_url"`
	GoalAmmount     int    `json:"goal_ammount"`
	Perks           string `json:"perks"`
	CurrentAmmount  int    `json:"current_ammount"`
	Slug            string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}

	campaignFormatter.Id = campaign.Id
	campaignFormatter.UserId = campaign.UserId
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShotDescription = campaign.ShortDescription
	campaignFormatter.Descrtiption = campaign.Description
	campaignFormatter.GoalAmmount = campaign.GoalAmmount
	campaignFormatter.Perks = campaign.Perks
	campaignFormatter.CurrentAmmount = campaign.CurrentAmmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImagesURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImagesURL = campaign.CampaignImages[0].FileName
	}
	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		CampaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, CampaignFormatter)

	}
	return campaignsFormatter
}

type FormatterCampaignDetail struct {
	Id               int                       `json:"id"`
	Name             string                    `json:"name"`
	UserId           int                       `json:"user_id"`
	ShortDescription string                    `json:"short_description"`
	Description      string                    `json:"description"`
	ImageURL         string                    `json:"image_url"`
	GoalAmmount      int                       `json:"goal_ammount"`
	CurrentAmmount   int                       `json:"current_ammount"`
	Slug             string                    `json:"slug"`
	Perks            []string                  `json:"perks"`
	User             UserCampaignFormatter     `json:"user"`
	Images           []ImagesCampaignFormatter `json:"images"`
}

type UserCampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type ImagesCampaignFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatDetailCampaign(campaign Campaign) FormatterCampaignDetail {
	campaignDetailFormatter := FormatterCampaignDetail{}
	campaignDetailFormatter.Id = campaign.Id
	campaignDetailFormatter.UserId = campaign.UserId
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmmount = campaign.GoalAmmount
	campaignDetailFormatter.CurrentAmmount = campaign.CurrentAmmount
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := UserCampaignFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName

	campaignDetailFormatter.User = campaignUserFormatter

	images := []ImagesCampaignFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := ImagesCampaignFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
