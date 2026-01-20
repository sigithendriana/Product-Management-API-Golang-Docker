package main

import (
	"my-product-api/internal/database"
	"my-product-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	r.POST("/products", handlers.CreateProduct)
	r.GET("/products", handlers.GetProducts)

	r.Run(":8080")
}
