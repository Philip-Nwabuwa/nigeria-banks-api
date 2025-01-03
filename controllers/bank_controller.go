package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nigeria-banks-api/database"
	"github.com/nigeria-banks-api/models"
)

func GetBanks(c *gin.Context) {
	count, err := database.GetBankCount()
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

	banks, err := database.GetAllBanks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching banks"})
		return
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

	if err := database.AddBank(&bank); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding bank"})
		return
	}

	c.JSON(http.StatusCreated, bank)
}
