package main

import (
	"api-golang/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	ProductController := controller.NewProductController()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
		})
	})

	server.GET("/products", ProductController.GetProduct)
	server.Run(":8080")
}
