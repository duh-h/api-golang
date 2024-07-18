package controller

import (
	"api-golang/usercase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsercase usercase.ProductUsercase
}

func NewProductController(usecase usercase.ProductUsercase) ProductController {
	return ProductController{
		ProductUsercase: usecase,
	}
}

func (p *ProductController) GetProduct(ctx *gin.Context) {

	products, err := p.ProductUsercase.GetProduct()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
