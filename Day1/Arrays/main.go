package main

import (
    "github.com/gin-gonic/gin"
    "retailer-api/config"
    "retailer-api/routes"
)

func main() {
    config.ConnectDatabase()
    r := gin.Default()
    routes.RegisterRoutes(r)
    r.Run()
}
