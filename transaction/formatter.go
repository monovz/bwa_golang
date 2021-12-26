package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	formattedTransactions := []CampaignTransactionFormatter{}
	for _, transaction := range transactions {
		formattedTransaction := FormatCampaignTransaction(transaction)
		formattedTransactions = append(formattedTransactions, formattedTransaction)
	}
	return formattedTransactions
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formattedTransaction := CampaignTransactionFormatter{}
	formattedTransaction.ID = transaction.ID
	formattedTransaction.Amount = transaction.Amount
	formattedTransaction.CreatedAt = transaction.CreatedAt
	formattedTransaction.Name = transaction.User.Name
	return formattedTransaction
}
