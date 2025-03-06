package controllers

import (
    "net/http"
    "retailer_project/config"
    "retailer_project/models"

    "github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&product)
    c.JSON(http.StatusOK, product)
}
