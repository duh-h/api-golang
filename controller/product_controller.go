package controller

import (
	"api-golang/model"
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

	products := []model.Product{
		{
			ID:    1,
			Name:  "batata",
			Price: 2.7,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
