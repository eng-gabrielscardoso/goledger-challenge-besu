package models

import (
	"fmt"
	"math/big"
	"time"
)

type Transaction struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Value     string    `gorm:"type:numeric;not null" json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Transaction) SetValue(value *big.Int) {
	t.Value = value.String()
}

func (t *Transaction) GetValue() (*big.Int, error) {
	value, success := new(big.Int).SetString(t.Value, 10)

	if !success {
		return nil, fmt.Errorf("failed to convert value from string to big.Int")
	}

	return value, nil
}
