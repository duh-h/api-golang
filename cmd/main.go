package main

import (
	"api-golang/controller"
	"api-golang/db"
	"api-golang/repository"
	"api-golang/server/middleware"
	"api-golang/service"
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

	LoginRepository := repository.NewLoginRepository(dbConnection)
	jwtService := service.NewJWTService()
	LoginUsercase := usercase.NewLoginUsecase(LoginRepository, &jwtService)
	LoginController := controller.NewLoginController(LoginUsercase)

	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
		})
	})

	server.POST("/login", LoginController.Login)
	server.POST("/users", UserController.CreateUser)

	productRoutes := server.Group("/product", middleware.Auth(jwtService))

	productRoutes.GET("/", ProductController.GetProduct)
	productRoutes.POST("/", ProductController.CreateProduct)
	productRoutes.GET("/:productId", ProductController.GetProductById)
	productRoutes.DELETE("/:productId", ProductController.DeleteProductById)
	productRoutes.PUT("/:productId", ProductController.UpdateProductById)

	server.Run(":8000")
}
