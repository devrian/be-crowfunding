package transaction

import (
	"be-crowfunding/user"
)

// GetCampaignTransactionsInput is ...
type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

// CreateTransactionInput is ...
type CreateTransactionInput struct {
	Amount     float64 `json:"amount" binding:"required"`
	CampaignID int     `json:"campaign_id" binding:"required"`
	User       user.User
}

// TransactionNotificationInput is ...
type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
