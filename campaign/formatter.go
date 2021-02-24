package campaign

import "strings"

// CampaignFormatter is ...
type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

// CampaignDetailFormatter is ...
type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	UserID           int                      `json:"user_id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	Slug             string                   `json:"slug"`
	Description      string                   `json:"description"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

// CampaignUserFormatter is ...
type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

// CampaignImageFormatter is ...
type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

// FormatCampaign is ...
func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaigFormatter := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		campaigFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaigFormatter
}

// FormatCampaigns is ...
func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	if len(campaigns) > 0 {
		for _, campaign := range campaigns {
			campaignFormatter := FormatCampaign(campaign)
			campaignsFormatter = append(campaignsFormatter, campaignFormatter)
		}
	}

	return campaignsFormatter
}

// FormatCampaignDetail is ...
func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
		Description:      campaign.Description,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	// mapping perks
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ", ") {
		perks = append(perks, perk)
	}
	campaignDetailFormatter.Perks = perks

	// mapping user
	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{
		Name:     user.Name,
		ImageURL: user.AvatarFileName,
	}
	campaignDetailFormatter.User = campaignUserFormatter

	// mapping images
	images := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{
			ImageURL:  image.FileName,
			IsPrimary: image.IsPrimary == 1,
		}
		images = append(images, campaignImageFormatter)
	}
	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
