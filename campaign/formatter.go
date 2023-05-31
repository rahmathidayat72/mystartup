package campaign

type CampaignFormatter struct {
	Id              int    `json: "id"`
	UserId          int    `json= "user_id"`
	Name            string `json: "name"`
	ShotDescription string `json: "short_description"`
	Descrtiption    string `json:"description"`
	ImagesURL       string `json:"images_url"`
	GoalAmmount     int    `json:"goal_ammount"`
	Perks           string `json: "perks"`
	CurrentAmmount  int    `json:"current_ammount"`
	Slug            string `json:"clug"`
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
