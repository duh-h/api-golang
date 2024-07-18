package main

import (
	"api-golang/controller"
	"api-golang/db"
	"api-golang/repository"
	"api-golang/usercase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsercase := usercase.NewProductUsercase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsercase)

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
		})
	})

	server.GET("/products", ProductController.GetProduct)
	server.Run(":8080")
}
