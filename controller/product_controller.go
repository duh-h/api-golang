package controller

import (
	"api-golang/model"
	"api-golang/usercase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsercase usercase.ProductUsecase
}

func NewProductController(usecase usercase.ProductUsecase) ProductController {
	return ProductController{
		ProductUsercase: usecase,
	}
}

func (p *ProductController) GetProduct(ctx *gin.Context) {

	products, err := p.ProductUsercase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {

	var product model.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.ProductUsercase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)

}

func (p *ProductController) GetProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.ProductUsercase.GetProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if product == nil {
		response := model.Response{
			Message: "Produto nao foi encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return

	}

	ctx.JSON(http.StatusOK, product)
}

func (p *ProductController) DeleteProductById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.ProductUsercase.DeleteProductById(productId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if product == nil {
		response := model.Response{
			Message: "Produto nao foi encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return

	}

	ctx.JSON(http.StatusOK, product)

}

func (p *ProductController) UpdateProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		response := model.Response{
			Message: "Dados do produto inválidos",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	updatedProduct, err := p.ProductUsercase.UpdateProductById(product.Price, productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}
