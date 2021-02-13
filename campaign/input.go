package campaign

import "be-crowfunding/user"

// GetCampaignDetailInput is ...
type GetCampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

// CreateCampaignInput is ...
type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}
