package controllers

import (
    "net/http"
    "retailer_project/config"
    "retailer_project/models"
    "time"

    "github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&customer)
    c.JSON(http.StatusOK, customer)
}

func DeleteInactiveCustomers(c *gin.Context) {
    cutoffTime := time.Now().AddDate(0, -1, 0)
    config.DB.Where("last_order_time < ?", cutoffTime).Delete(&models.Customer{})
    c.JSON(http.StatusOK, gin.H{"message": "Inactive customers deleted"})
}
