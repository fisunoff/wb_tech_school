package usecase

import (
	"errors"
	"order_service/internal/model"
	"order_service/internal/repository"
)

var ErrEmptyUid = errors.New("uid is empty")
var ErrOrderNotFound = errors.New("order not found")

// OrderUseCase реализует бизнес-логику работы с заказами.
type OrderUseCase struct {
	repo repository.OrderRepository
}

func NewOrderUseCase(orderDb repository.OrderRepository) *OrderUseCase {
	return &OrderUseCase{repo: orderDb}
}

// GetOrderByUID возвращает заказ по уникальному идентификатору.
// Возвращает ErrEmptyUid, если UID пустой.
// Возвращает ErrOrderNotFound, если заказ не найден.
func (uc *OrderUseCase) GetOrderByUID(uid string) (*model.Order, error) {
	if uid == "" {
		return nil, ErrEmptyUid
	}
	ans, err := uc.repo.FindByUID(uid)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	return ans, nil
}
