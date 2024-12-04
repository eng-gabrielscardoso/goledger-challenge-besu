package services

import (
	"fmt"
	"math/big"

	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/config"
	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/models"
	"github.com/eng-gabrielscardoso/goledger-challenge-besu/internal/services/blockchain"
	"gorm.io/gorm"
)

type TransactionService struct{}

func New() *TransactionService {
	return &TransactionService{}
}

func (transactionService *TransactionService) GetValue() (*big.Int, error) {
	value := blockchain.CallContract("get")

	if value == nil {
		return nil, fmt.Errorf("failed to retrieve value from contract")
	}

	return value, nil
}

func (transactionService *TransactionService) SetValue(value *big.Int) error {
	if blockchain.ExecContract("set", value); value == nil {
		return fmt.Errorf("failed to set a new value for contract")
	}

	return nil
}

func (transactionService *TransactionService) SyncTransaction() error {
	value, err := transactionService.GetValue()

	if err != nil {
		return err
	}

	var existingTransaction models.Transaction

	if err := config.GetDatabaseConnection().First(&existingTransaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			transaction := &models.Transaction{}
			transaction.SetValue(value)

			if err := config.GetDatabaseConnection().Create(transaction).Commit().Error; err != nil {
				return fmt.Errorf("failed to create new transaction: %v", err)
			}

			return nil
		}

		return fmt.Errorf("failed to check transaction: %v", err)
	}

	existingTransaction.SetValue(value)

	if err := config.GetDatabaseConnection().Model(&existingTransaction).Update("value", value.String()).Error; err != nil {
		return fmt.Errorf("failed to update existing transaction: %v", err)
	}

	return nil
}

func (transactionService *TransactionService) CheckTransaction() (bool, error) {
	value, err := transactionService.GetValue()

	if err != nil {
		return false, err
	}

	var transaction models.Transaction

	if err := config.GetDatabaseConnection().First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, fmt.Errorf("no transaction found")
		}

		return false, fmt.Errorf("failed to check transaction: %v", err)
	}

	transactionValue, err := transaction.GetValue()

	if err != nil {
		return false, err
	}

	if transactionValue.Cmp(value) == 0 {
		return true, nil
	}

	return false, nil
}
