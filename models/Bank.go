package model

import (
	"app_backend/database"

	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	Name       string    `gorm:"size:100;not null;" json:"name"`
	Type       string    `gorm:"size:100;not null;" json:"type"`
	Account    string    `gorm:"size:100;not null;" json:"account"`
	Number     string    `gorm:"size:50;not null;" json:"number"`
	Total      float64   `gorm:"not null;" json:"total"`
	Products   []Product `gorm:"foreignKey:ProductID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CustomerID uint      `json:"customer_id"`
	Customer   Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (bank *Bank) Save() (*Bank, error) {
	err := database.Database.Create(&bank).Error
	if err != nil {
		return &Bank{}, err
	}
	return bank, nil
}
