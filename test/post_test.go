package test

import (
	"api-golang/controller"
	"api-golang/db"
	"api-golang/model"
	"api-golang/repository"
	"api-golang/usercase"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPOSTproduct(t *testing.T) {

	product := model.Product{
		Name:  "Product test",
		Price: 11.1,
	}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(product)

	if err != nil {
		t.Error(err)
	}

	dbConnection, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConnection.Close()

	router := gin.Default()

	productRepository := repository.NewProductRepository(dbConnection)
	productUsecase := usercase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUsecase)

	router.POST("/product", productController.CreateProduct)
	router.PUT("/product/:productId", productController.UpdateProductById)
	router.DELETE("/product/:productId", productController.DeleteProductById)

	req := httptest.NewRequest(http.MethodPost, "/product", &b)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Expected status code 201 Created")

	var response model.Product
	err = json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	assert.Equal(t, product.Name, response.Name, "Expected product name to match")
	assert.Equal(t, product.Price, response.Price, "Expected product price to match")
	assert.NotZero(t, response.ID, "Expected product ID to be non-zero")

	productPUT := model.Product{
		Price: 1,
	}

	var bPUT bytes.Buffer
	err = json.NewEncoder(&bPUT).Encode(productPUT)
	if err != nil {
		t.Error(err)
	}

	urlPUT := fmt.Sprintf("/product/%d", response.ID)
	putReq := httptest.NewRequest(http.MethodPut, urlPUT, &bPUT)
	putReq.Header.Set("Content-Type", "application/json")
	putW := httptest.NewRecorder()
	router.ServeHTTP(putW, putReq)

	assert.Equal(t, http.StatusOK, putW.Code, "Expected status code 200 OK")

	var responsePUT model.Product
	err = json.NewDecoder(putW.Body).Decode(&responsePUT)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	assert.Equal(t, productPUT.Price, responsePUT.Price, "Expected product price to match")

	url := fmt.Sprintf("/product/%d", response.ID)

	deleteReq := httptest.NewRequest(http.MethodDelete, url, nil)
	deleteW := httptest.NewRecorder()
	router.ServeHTTP(deleteW, deleteReq)

	assert.Equal(t, http.StatusOK, deleteW.Code, "Expected status code 200 OK")

}
