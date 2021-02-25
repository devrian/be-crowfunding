package transaction

import (
	"be-crowfunding/campaign"
	"be-crowfunding/user"
	"time"
)

// Transaction is ...
type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     float64
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
