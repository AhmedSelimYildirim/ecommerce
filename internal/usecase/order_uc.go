package usecase

import "github.com/AhmedSelimYildirim/ecommerce/internal/domain/models"

type OrderUsecase struct {
	// order repository burada olacak
}

func NewOrderUsecase() *OrderUsecase {
	return &OrderUsecase{}
}

func (uc *OrderUsecase) CreateOrder(order *models.Order) error {
	// TODO: implement order creation
	return nil
}

func (uc *OrderUsecase) GetOrdersByUser(userID int) ([]models.Order, error) {
	// TODO: implement fetching user orders
	return nil, nil
}
