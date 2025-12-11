package pg

import (
	"github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id int) (*models.User, error)
}
