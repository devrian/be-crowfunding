package transaction

import "time"

// CampaignTransactionsFormatter is ...
type CampaignTransactionsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// UserTransactionFormatter is ...
type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    float64           `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

// CampaignFormatter is ...
type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

// TransactionFormatter is ...
type TransactionFormatter struct {
	ID         int       `json:"id"`
	CampaignID int       `json:"campaign_id"`
	UserID     int       `json:"user_id"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	PaymentURL string    `json:"payment_url"`
	CreatedAt  time.Time `json:"created_at"`
}

// FormatCampaignTransaction is ...
func FormatCampaignTransaction(transaction Transaction) CampaignTransactionsFormatter {
	formatter := CampaignTransactionsFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatter
}

// FormatCampaignTransactions is ...
func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionsFormatter {
	transactionsFormatter := []CampaignTransactionsFormatter{}

	if len(transactions) > 0 {
		for _, transaction := range transactions {
			transactionFormatter := FormatCampaignTransaction(transaction)
			transactionsFormatter = append(transactionsFormatter, transactionFormatter)
		}
	}

	return transactionsFormatter
}

// FormatUserTransaction is ...
func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
	}

	campaignFormatter := CampaignFormatter{
		Name: transaction.Campaign.Name,
	}

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

// FormatUserTransactions is ...
func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	transactionsFormatter := []UserTransactionFormatter{}

	if len(transactions) > 0 {
		for _, transaction := range transactions {
			transactionFormatter := FormatUserTransaction(transaction)
			transactionsFormatter = append(transactionsFormatter, transactionFormatter)
		}
	}

	return transactionsFormatter
}

// FormatTransaction is ...
func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:         transaction.ID,
		CampaignID: transaction.CampaignID,
		UserID:     transaction.UserID,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Code,
		PaymentURL: transaction.PaymentURL,
		CreatedAt:  transaction.CreatedAt,
	}

	return formatter
}
