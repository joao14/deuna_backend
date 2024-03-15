package model

import (
	"app_backend/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	Name       string   `gorm:"size:100;not null;" json:"name"`
	Type       string   `gorm:"size:100;not null;" json:"type"`
	Account    string   `gorm:"size:100;not null;" json:"account"`
	Number     string   `gorm:"size:50;not null;" json:"number"`
	Total      float64  `gorm:"not null;" json:"total"`
	CustomerID uint     `json:"customer_id"`
	Customer   Customer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (bank *Bank) Save() (*Bank, error) {
	accountHash, err := bcrypt.GenerateFromPassword([]byte(bank.Account), bcrypt.DefaultCost)
	if err != nil {
		return &Bank{}, err
	}

	bank.Account = string(accountHash)

	err = database.Database.Create(&bank).Error
	if err != nil {
		return &Bank{}, err
	}
	return bank, nil
}
