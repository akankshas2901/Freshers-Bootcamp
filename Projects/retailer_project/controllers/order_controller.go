package controllers

import (
    "net/http"
    "retailer_project/config"
    "retailer_project/models"
    "time"

    "github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    order.CreatedAt = time.Now()
    order.Status = "Placed"
    config.DB.Create(&order)

    config.DB.Model(&models.Customer{}).Where("id = ?", order.CustomerID).
        Update("last_order_time", time.Now())

    c.JSON(http.StatusOK, order)
}
