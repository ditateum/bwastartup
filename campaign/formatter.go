package campaign

type CampaignFormater struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormater {
	campaignFormater := CampaignFormater{}
	campaignFormater.ID = campaign.ID
	campaignFormater.UserID = campaign.UserID
	campaignFormater.ShortDescription = campaign.ShortDescription
	campaignFormater.GoalAmount = campaign.GoalAmount
	campaignFormater.CurrentAmount = campaign.CurrentAmount
	campaignFormater.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormater.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormater
}

func FormatCampaigns(campaign []Campaign) []CampaignFormater {
	campaignsFormatter := []CampaignFormater{}

	for _, campaign := range campaign {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
