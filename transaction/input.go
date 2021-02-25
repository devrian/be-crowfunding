package transaction

import "be-crowfunding/user"

// GetCampaignTransactionsInput is ...
type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
