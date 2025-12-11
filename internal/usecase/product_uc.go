package usecase

import (
	"github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"
	"github.com/AhmedSelimYildirim/ecommerce/internal/repository/pg"
)

type ProductUsecase struct {
	repo pg.ProductRepository
}

func NewProductUsecase(repo pg.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: repo}
}

func (uc *ProductUsecase) CreateProduct(p *models.Product) error {
	return uc.repo.Create(p)
}

func (uc *ProductUsecase) GetProductByID(id int) (*models.Product, error) {
	return uc.repo.GetByID(id)
}

func (uc *ProductUsecase) UpdateProduct(p *models.Product) error {
	return uc.repo.Update(p)
}

func (uc *ProductUsecase) DeleteProduct(id int) error {
	return uc.repo.Delete(id)
}

func (uc *ProductUsecase) GetAllProducts() ([]*models.Product, error) {
	return uc.repo.GetAll()
}
