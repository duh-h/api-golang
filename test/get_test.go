package test

import (
	"api-golang/controller"
	"api-golang/db"
	"api-golang/repository"
	"api-golang/usercase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGETProduct(t *testing.T) {

	dbConnection, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConnection.Close()

	router := gin.Default()

	productRepository := repository.NewProductRepository(dbConnection)
	productUsecase := usercase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUsecase)

	router.GET("/product", productController.GetProduct)

	req := httptest.NewRequest(http.MethodGet, "/product", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200 OK")

}
