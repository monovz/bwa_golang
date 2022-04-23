package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
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

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	formattedTransactions := []UserTransactionFormatter{}
	for _, transaction := range transactions {
		formattedTransaction := FormatUserTransaction(transaction)
		formattedTransactions = append(formattedTransactions, formattedTransaction)
	}
	return formattedTransactions
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formattedTransaction := UserTransactionFormatter{}
	formattedTransaction.ID = transaction.ID
	formattedTransaction.Amount = transaction.Amount
	formattedTransaction.Status = transaction.Status
	formattedTransaction.CreatedAt = transaction.CreatedAt
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].Filename
	}
	formattedTransaction.Campaign = campaignFormatter
	return formattedTransaction
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formattedTransaction := TransactionFormatter{}
	formattedTransaction.ID = transaction.ID
	formattedTransaction.CampaignID = transaction.CampaignID
	formattedTransaction.UserID = transaction.UserID
	formattedTransaction.Status = transaction.Status
	formattedTransaction.Code = transaction.Code
	formattedTransaction.Amount = transaction.Amount
	formattedTransaction.PaymentURL = transaction.PaymentURL
	return formattedTransaction
}
