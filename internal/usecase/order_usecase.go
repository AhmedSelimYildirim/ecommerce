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

// CreateOrder yeni sipariş oluşturur
func (uc *OrderUsecase) CreateOrder(userID int, productIDs []int) error {
	return uc.orderRepo.Create(userID, productIDs)
}

// GetUserOrders kullanıcıya ait siparişleri döner
func (uc *OrderUsecase) GetUserOrders(userID int) ([]pg.Order, error) {
	return uc.orderRepo.GetByUserID(userID)
}
