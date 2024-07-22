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
			Message: "Id so produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if id == "" {
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
