package usecase

import (
	"context"
	"errors"
	"order_service/internal/model"
	"order_service/internal/repository"
)

type OrderCreateUseCase struct {
	repo repository.OrderRepository
}

func NewOrderCreateUseCase(repo repository.OrderRepository) *OrderCreateUseCase {
	return &OrderCreateUseCase{repo: repo}
}

// CreateOrderFromRaw создает заказ из сырых данных (например, из Kafka).
// Возвращает ошибку, если заказ невалиден или не удалось сохранить.
func (uc *OrderCreateUseCase) CreateOrderFromRaw(ctx context.Context, rawData []byte) error {
	order, err := model.ParseOrderFromJSON(rawData)
	if err != nil {
		return ErrInvalidOrderData
	}

	return uc.repo.Save(order)
}

var ErrInvalidOrderData = errors.New("invalid order data")
