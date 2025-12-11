package usecase

import (
	"errors"
	"time"

	"github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"
	"github.com/AhmedSelimYildirim/ecommerce/internal/repository/pg"
	"github.com/AhmedSelimYildirim/ecommerce/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo      pg.UserRepository
	jwtSecret string
	jwtExpire time.Duration
}

func NewUserUsecase(repo pg.UserRepository, jwtSecret string, jwtExpireHours int) *UserUsecase {
	return &UserUsecase{
		repo:      repo,
		jwtSecret: jwtSecret,
		jwtExpire: time.Duration(jwtExpireHours) * time.Hour,
	}
}

// RegisterUser creates a new user
func (uc *UserUsecase) RegisterUser(user *models.User) error {
	// Email kontrolü
	existing, _ := uc.repo.FindByEmail(user.Email)
	if existing != nil {
		return errors.New("email already exists")
	}

	// Şifre hashleme
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashed)
	user.Role = "user"
	user.CreatedAt = time.Now()

	return uc.repo.Create(user)
}

// LoginUser checks email & password, returns JWT
func (uc *UserUsecase) LoginUser(email, password string) (string, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	// Şifre kontrolü
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// JWT üret
	token, err := jwt.GenerateToken(user.ID, user.Role, uc.jwtSecret, uc.jwtExpire)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserByID returns user by ID
func (uc *UserUsecase) GetUserByID(id int) (*models.User, error) {
	return uc.repo.FindByID(id)
}
