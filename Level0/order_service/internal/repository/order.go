package repository

import (
	"errors"
	"order_service/internal/model"
)

var ErrNotFound = errors.New("not found")

type OrderRepository interface {
	Save(order *model.Order) error
	FindByUID(uid string) (*model.Order, error)
	FindTopNewest(limit int) ([]*model.Order, error)
}
