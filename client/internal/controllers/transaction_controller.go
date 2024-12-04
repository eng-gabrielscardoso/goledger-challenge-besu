package controllers

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/services"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service *services.TransactionService
}

type SetValueRequest struct {
	Value *big.Int `json:"value" binding:"required"`
}

func New(service *services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func TransactionResource(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (transactionController *TransactionController) GetValue(c *gin.Context) {
	value, err := transactionController.service.GetValue()
	TransactionResource(c, gin.H{"value": value}, err)
}

func (transactionController *TransactionController) SetValue(c *gin.Context) {
	var request SetValueRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		TransactionResource(c, nil, err)
		return
	}

	if request.Value == nil {
		TransactionResource(c, nil, fmt.Errorf("invalid value format"))
		return
	}

	err := transactionController.service.SetValue(request.Value)
	TransactionResource(c, gin.H{"message": "Value set successfully"}, err)
}

func (transactionController *TransactionController) SyncTransaction(c *gin.Context) {
	err := transactionController.service.SyncTransaction()
	TransactionResource(c, gin.H{"message": "Transaction synchronised successfully"}, err)
}

func (transactionController *TransactionController) CheckTransaction(c *gin.Context) {
	match, err := transactionController.service.CheckTransaction()
	TransactionResource(c, gin.H{"match": match}, err)
}
