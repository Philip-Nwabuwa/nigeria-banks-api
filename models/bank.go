package models

type Bank struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Code            string `json:"code"`
	USSDCode        string `json:"ussd_code"`
	BaseUSSDCode    string `json:"base_ussd_code"`
	BankCategory    string `json:"bank_category"`
	InternetBanking bool   `json:"internet_banking"`
}

type BanksResponse struct {
	Banks []Bank `json:"banks"`
}
