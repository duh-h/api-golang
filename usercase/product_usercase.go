package usercase

import (
	"api-golang/model"
	"api-golang/repository"
)

type ProductUsercase struct {
	repository repository.ProductRepository
}

func NewProductUsercase(repo repository.ProductRepository) ProductUsercase {
	return ProductUsercase{
		repository: repo,
	}
}

func (pu *ProductUsercase) GetProduct() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
