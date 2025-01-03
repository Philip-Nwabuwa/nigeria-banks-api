package models

import (
	"github.com/google/uuid"
)

type Bank struct {
	ID              string `json:"id"`
	Name            string `json:"name" binding:"required"`
	Code            string `json:"code" binding:"required"`
	USSDCode        string `json:"ussd_code" binding:"required"`
	BaseUSSDCode    string `json:"base_ussd_code" binding:"required"`
	BankCategory    string `json:"bank_category" binding:"required"`
	InternetBanking bool   `json:"internet_banking"`
}

type BanksData struct {
	Banks []Bank `json:"banks"`
}

type APIResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func NewAPIResponse(message string, status int, data interface{}) APIResponse {
	return APIResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
}

func (b *Bank) Validate() bool {
	return b.Name != "" && 
		   b.Code != "" && 
		   b.USSDCode != "" && 
		   b.BaseUSSDCode != "" && 
		   b.BankCategory != ""
}

func NewBank(name string, code string, ussdCode string, baseUssdCode string, bankCategory string, internetBanking bool) Bank {
	return Bank{
		ID:              uuid.New().String(),
		Name:            name,
		Code:            code,
		USSDCode:        ussdCode,
		BaseUSSDCode:    baseUssdCode,
		BankCategory:    bankCategory,
		InternetBanking: internetBanking,
	}
}
