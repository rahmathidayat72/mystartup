package transaction

import "time"

type CampaignTransactionFormatter struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Amonnt   int       `json:"amount"`
	CreateAt time.Time `json:"create_at"`
}

func FormatterCampaignTransaction(transaction Transactions) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.Id = transaction.Id
	formatter.Name = transaction.User.Name
	formatter.Amonnt = transaction.Amount
	formatter.CreateAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []Transactions) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatterCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)

	}
	return transactionsFormatter
}
