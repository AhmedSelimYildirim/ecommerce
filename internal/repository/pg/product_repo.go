package pg

import "github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"

type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id int) (*models.Product, error)
	Update(product *models.Product) error
	Delete(id int) error
	GetAll() ([]*models.Product, error)
}
