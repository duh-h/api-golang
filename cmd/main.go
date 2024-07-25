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
	ProductUsercase := usercase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsercase)

	UserRepository := repository.NewUserRepository(dbConnection)
	UserUsercase := usercase.NewUserUseCase(UserRepository)
	UserController := controller.NewUserController(UserUsercase)

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
		})
	})

	server.POST("/users", UserController.CreateUser)

	server.GET("/products", ProductController.GetProduct)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.DELETE("/product/:productId", ProductController.DeleteProductById)
	server.PUT("/product/:productId", ProductController.UpdateProductById)
	server.Run(":8000")
}
