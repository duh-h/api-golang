package usercase

import "api-golang/model"

type ProductUsercase struct {
}

func NewProductUsercase() ProductUsercase {
	return ProductUsercase{}
}

func (pu *ProductUsercase) GetProduct() ([]model.Product, error) {
	return []model.Product{}, nil
}
