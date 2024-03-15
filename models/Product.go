package model

import (
	"app_backend/database"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID uint
	Name      string `gorm:"size:100;not null;" json:"name"`
	Stock     int    `gorm:"size:4;not null;" json:"stock"`
	Color     string `gorm:"size:100;not null;" json:"color"`
}

func (product *Product) Save() (*Product, error) {
	err := database.Database.Create(&product).Error
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}
