package transaction

import (
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
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
