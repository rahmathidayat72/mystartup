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

type UserTransactionFormatter struct {
	Id       int               `json:"id"`
	Amount   int               `json:"amount"`
	Status   string            `json:"starus"`
	CreateAt time.Time         `json:"create_at"`
	Campaign campaignFormatter `json:"campaign"`
}

type campaignFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatUserTransacstion(transaction Transactions) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.Id = transaction.Id
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreateAt = transaction.CreatedAt

	campaignFormatter := campaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageUrl = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter

}

func FormatUserTransactions(transactions []Transactions) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransacstion(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)

	}
	return transactionsFormatter
}
