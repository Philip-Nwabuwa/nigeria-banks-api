package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nigeria-banks-api/database"
	"github.com/nigeria-banks-api/models"
)

func GetBanks(c *gin.Context) {
	count, err := database.GetBankCount()
	if err != nil {
		response := models.NewAPIResponse("Error checking bank count", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if count == 0 {
		emptyResponse := models.NewAPIResponse("No banks found", http.StatusOK, models.BanksData{Banks: []models.Bank{}})
		c.JSON(http.StatusOK, emptyResponse)
		return
	}

	banks, err := database.GetAllBanks()
	if err != nil {
		response := models.NewAPIResponse("Error fetching banks", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	banksData := models.BanksData{
		Banks: banks,
	}

	response := models.NewAPIResponse("Banks fetched successfully", http.StatusOK, banksData)
	c.JSON(http.StatusOK, response)
}

func AddBank(c *gin.Context) {
	var bank models.Bank
	if err := c.ShouldBindJSON(&bank); err != nil {
		response := models.NewAPIResponse("Invalid request body", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if !bank.Validate() {
		response := models.NewAPIResponse("All fields are required", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := database.AddBank(&bank); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			response := models.NewAPIResponse(err.Error(), http.StatusConflict, nil)
			c.JSON(http.StatusConflict, response)
			return
		}
		response := models.NewAPIResponse("Error adding bank", http.StatusInternalServerError, nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := models.NewAPIResponse("Bank added successfully", http.StatusCreated, bank)
	c.JSON(http.StatusCreated, response)
}
