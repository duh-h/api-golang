package controller

import (
	"api-golang/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func NewProductController() ProductController {
	return ProductController{}
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
