package campaign

import "be-crowfunding/user"

// GetCampaignDetailInput is ...
type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

// CreateCampaignInput is ...
type CreateCampaignInput struct {
	Name             string  `json:"name" binding:"required"`
	ShortDescription string  `json:"short_description" binding:"required"`
	Description      string  `json:"description" binding:"required"`
	GoalAmount       float64 `json:"goal_amount" binding:"required"`
	Perks            string  `json:"perks" binding:"required"`
	User             user.User
}

// CreateCampaignImageInput is ...
type CreateCampaignImageInput struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}
