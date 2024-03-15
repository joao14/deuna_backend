package model

import (
	"app_backend/database"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Identify string `gorm:"size:13;not null;unique" json:"identify"`
	Name     string `gorm:"size:255;not null;" json:"name"`
	Lastname string `gorm:"size:255;not null;" json:"lastname"`
	Age      string `gorm:"size:3;not null;" json:"age"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Country  string `gorm:"size:100;not null;" json:"country"`
}

func (customer *Customer) Save() (*Customer, error) {
	err := database.Database.Create(&customer).Error
	if err != nil {
		return &Customer{}, err
	}
	return customer, nil
}

func FindCustonerByIdentify(identify string) (Customer, error) {
	var customer Customer
	err := database.Database.Preload("Customers").Where("identify=?", identify).Find(&customer).Error
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}

func (customer *Customer) FindManyCustomer() ([]Customer, error) {
	var customers []Customer
	err := database.Database.Preload("Customers").Find(&customers).Error
	if err != nil {
		return []Customer{}, err
	}
	return customers, nil
}
