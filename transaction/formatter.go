package transaction

import "time"

// CampaignTransactionsFormatter is ...
type CampaignTransactionsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
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
