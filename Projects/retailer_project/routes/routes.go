package routes

import (
    "github.com/gin-gonic/gin"
    "retailer_project/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/customers", controllers.CreateCustomer)
    r.DELETE("/customers/inactive", controllers.DeleteInactiveCustomers)

    r.POST("/products", controllers.CreateProduct)

    r.POST("/orders", controllers.PlaceOrder)

    return r
}
