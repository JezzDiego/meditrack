package usecase

import (
	"meditrack/model"
	"meditrack/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (p *ProductUsecase) GetAllProducts() ([]model.FullProduct, error) {
	products, err := p.repository.GetAllProducts()
	if err != nil {
		return []model.FullProduct{}, err
	}

	return products, nil
}

func (p *ProductUsecase) GetProductById(id int) (*model.FullProduct, error) {
	product, err := p.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductUsecase) GetProductByGtin(gtin string) (*model.FullProduct, error) {
	product, err := p.repository.GetProductByGtin(gtin)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductUsecase) CreateProduct(product model.FullProduct) (model.FullProduct, error) {
	productId, err := p.repository.CreateProduct(product)
	if err != nil {
		return model.FullProduct{}, err
	}

	product.ID = productId

	return product, nil
}
