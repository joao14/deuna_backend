package controller

import (
	model "app_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProduct(context *gin.Context) {
	var input model.Product

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
