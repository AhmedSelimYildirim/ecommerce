package usecase

import (
	"github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"
	"github.com/AhmedSelimYildirim/ecommerce/internal/repository/pg"
)

type CartUsecase struct {
	repo pg.CartRepository
}

func NewCartUsecase(repo pg.CartRepository) *CartUsecase {
	return &CartUsecase{repo: repo}
}

func (uc *CartUsecase) AddItem(item *models.CartItem) error {
	return uc.repo.AddItem(item)
}

func (uc *CartUsecase) UpdateItem(item *models.CartItem) error {
	return uc.repo.UpdateItem(item)
}

func (uc *CartUsecase) DeleteItem(id int) error {
	return uc.repo.DeleteItem(id)
}

func (uc *CartUsecase) GetCart(userID int) ([]*models.CartItem, error) {
	return uc.repo.GetCartByUser(userID)
}

func (uc *CartUsecase) ClearCart(userID int) error {
	return uc.repo.ClearCart(userID)
}
