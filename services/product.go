package services

import (
	"errors"
	productdto "testgo/dto/product"
	"testgo/models"

	"github.com/go-playground/validator/v10"
)

func (s *ServicessInit) ListProductsPaginated(page, limit int, search string) (products []models.Product, totalCount int, err error) {
	products, totalCount, err = s.RepoProduct.ListProductsPaginated(page, limit, search)
	if err != nil {
		return nil, 0, err
	}
	return products, totalCount, nil
}

func (s *ServicessInit) GetProductByID(id int) (models.Product, error) {
	product, err := s.RepoProduct.GetProduct(id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (s *ServicessInit) CreateProduct(req productdto.CreateProductRequest) (models.Product, error) {

	product := models.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Image:       req.Image,
	}

	validation := validator.New()
	if err := validation.Struct(req); err != nil {
		return models.Product{}, err
	}

	createdProduct, err := s.RepoProduct.CreateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	return createdProduct, nil
}

func (s *ServicessInit) UpdateProduct(id int, isAdmin bool, request productdto.Update1ProductRequest) (models.Product, error) {
	if !isAdmin {
		return models.Product{}, errors.New("sorry you're not admin")
	}

	validation := validator.New()
	if err := validation.Struct(request); err != nil {
		return models.Product{}, err
	}

	product, err := s.RepoProduct.GetProduct(id)
	if err != nil {
		return models.Product{}, err
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	updatedProduct, err := s.RepoProduct.UpdateProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	return updatedProduct, nil
}

func (s *ServicessInit) DeleteProduct(id int, isAdmin bool) (models.Product, error) {
	if !isAdmin {
		return models.Product{}, errors.New("sorry you're not admin")
	}

	product, err := s.RepoProduct.GetProduct(id)
	if err != nil {
		return models.Product{}, err
	}

	deletedProduct, err := s.RepoProduct.DeleteProduct(product)
	if err != nil {
		return models.Product{}, err
	}

	return deletedProduct, nil
}
