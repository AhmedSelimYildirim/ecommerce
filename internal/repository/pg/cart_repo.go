package pg

import "github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"

type CartRepository interface {
	AddItem(item *models.CartItem) error
	UpdateItem(item *models.CartItem) error
	DeleteItem(id int) error
	GetCartByUser(userID int) ([]*models.CartItem, error)
	ClearCart(userID int) error
}
