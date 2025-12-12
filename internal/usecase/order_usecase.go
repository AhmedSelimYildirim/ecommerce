package usecase

import "github.com/AhmedSelimYildirim/ecommerce/internal/repository/pg"

type OrderUsecase struct {
	orderRepo pg.OrderRepository
}

func NewOrderUsecase(orderRepo pg.OrderRepository) *OrderUsecase {
	return &OrderUsecase{
		orderRepo: orderRepo,
	}
}

func (uc *OrderUsecase) CreateOrder(userID int, productIDs []int) error {
	// TODO: implement logic
	return uc.orderRepo.Create(userID, productIDs)
}

func (uc *OrderUsecase) GetUserOrders(userID int) ([]pg.Order, error) {
	// TODO: implement logic
	return uc.orderRepo.GetByUserID(userID)
}
