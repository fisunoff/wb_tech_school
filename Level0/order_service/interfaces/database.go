package interfaces

import (
	"github.com/jmoiron/sqlx"
	"order_service/internal/model"
)

type DatabaseInterface interface {
	Close() error
	GetDb() *sqlx.DB
	Save(order *model.Order) error
	GetOrderByUID(orderUID string) (*model.Order, error)
	GetTopNewestOrders(quantity int) ([]*model.Order, error)
}
