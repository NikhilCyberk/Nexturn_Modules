package service

import (
	"Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/model"
	"Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/repository"
)

type ProductService struct {
    repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *model.Product) error {
    return s.repo.Create(product)
}

func (s *ProductService) GetProduct(id int) (*model.Product, error) {
    return s.repo.GetByID(id)
}

func (s *ProductService) GetProducts(page, limit int) ([]model.Product, int, error) {
    return s.repo.GetAll(page, limit)
}

func (s *ProductService) UpdateProduct(product *model.Product) error {
    return s.repo.Update(product)
}

func (s *ProductService) UpdateStock(id, stock int) error {
    return s.repo.UpdateStock(id, stock)
}

func (s *ProductService) DeleteProduct(id int) error {
    return s.repo.Delete(id)
}