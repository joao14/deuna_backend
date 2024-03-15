package controller

import (
	"app_backend/helper"
	model "app_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCustomer(context *gin.Context) {
	var input model.Customer

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedCustomer, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedCustomer})
}

func GetOneCustomer(context *gin.Context) {
	customers, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customers.Entries})
}

func GetAllCustomer(context *gin.Context) {
	var input model.Customer
	print("DATA....")
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customers, err := input.FindManyCustomer()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customers})
}
