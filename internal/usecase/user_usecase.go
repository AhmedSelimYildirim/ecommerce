package usecase

import (
	"errors"
	"github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"
	"github.com/AhmedSelimYildirim/ecommerce/internal/repository/pg"
)

type UserUsecase struct {
	repo       pg.UserRepository
	jwtSecret  string
	jwtExpires int
}

func NewUserUsecase(repo pg.UserRepository, jwtSecret string, jwtExpires int) *UserUsecase {
	return &UserUsecase{
		repo:       repo,
		jwtSecret:  jwtSecret,
		jwtExpires: jwtExpires,
	}
}

// RegisterUser artık (createdUser, error) dönecek
func (uc *UserUsecase) RegisterUser(user *models.User) (*models.User, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return nil, errors.New("missing fields")
	}

	createdUser, err := uc.repo.Create(user) // repo.Create User döndürmeli
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// LoginUser
func (uc *UserUsecase) LoginUser(email, password string) (string, error) {
	// JWT oluşturma vs.
	return "mock-jwt-token", nil
}

// GetUserByID
func (uc *UserUsecase) GetUserByID(id int) (*models.User, error) {
	return uc.repo.FindByID(id) // <- burayı FindByID yap
}
