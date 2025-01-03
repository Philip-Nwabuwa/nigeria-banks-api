package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nigeria-banks-api/database"
	"github.com/nigeria-banks-api/models"
)

func GetBanks(c *gin.Context) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM banks").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking bank count"})
		return
	}

	if count == 0 {
		response := models.BanksResponse{
			Banks: []models.Bank{},
		}
		c.JSON(http.StatusOK, response)
		return
	}

	rows, err := database.DB.Query("SELECT id, name, code, ussd_code, base_ussd_code, bank_category, internet_banking FROM banks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching banks"})
		return
	}
	defer rows.Close()

	var banks []models.Bank
	for rows.Next() {
		var bank models.Bank
		err := rows.Scan(&bank.ID, &bank.Name, &bank.Code, &bank.USSDCode, &bank.BaseUSSDCode, &bank.BankCategory, &bank.InternetBanking)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning bank data"})
			return
		}
		banks = append(banks, bank)
	}

	response := models.BanksResponse{
		Banks: banks,
	}

	c.JSON(http.StatusOK, response)
}

func AddBank(c *gin.Context) {
	var bank models.Bank
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if bank.Name == "" || bank.Code == "" || bank.USSDCode == "" || bank.BaseUSSDCode == "" || bank.BankCategory == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	stmt, err := database.DB.Prepare(`
		INSERT INTO banks (name, code, ussd_code, base_ussd_code, bank_category, internet_banking) 
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparing statement"})
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(bank.Name, bank.Code, bank.USSDCode, bank.BaseUSSDCode, bank.BankCategory, bank.InternetBanking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting bank data"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting inserted ID"})
		return
	}

	bank.ID = int(id)
	c.JSON(http.StatusCreated, bank)
}
