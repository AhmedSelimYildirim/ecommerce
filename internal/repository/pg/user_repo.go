package pg

import (
	"github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"
)

// UserRepository interface’i artık Create fonksiyonu ile birlikte createdUser döndürüyor
type UserRepository interface {
	Create(user *models.User) (*models.User, error) // <- düzeltildi
	FindByEmail(email string) (*models.User, error)
	FindByID(id int) (*models.User, error)
}
