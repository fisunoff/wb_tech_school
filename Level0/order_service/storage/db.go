package storage

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"order_service/model"
)

var defaultCacheSizeError = errors.New("defaultCacheSize must be greater than zero")
var startCacheSizeError = errors.New("startCacheSize must be greater than zero")

type Storage struct {
	db    *sqlx.DB
	cache map[string]*model.Order
}

// New - подключение к базе.
func New(connStr string, defaultCacheSize int, startCacheSize int) (*Storage, error) {
	if defaultCacheSize < 1 {
		return nil, defaultCacheSizeError
	}
	if startCacheSize < 1 {
		return nil, startCacheSizeError
	}

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}
	storage := &Storage{
		db:    db,
		cache: make(map[string]*model.Order),
	}
	orders, err := storage.GetTopNewestOrders(startCacheSize)
	if err != nil {
		return nil, fmt.Errorf("ошибка при заполнении кэша: %w", err)
	}
	for _, order := range orders {
		storage.cache[order.OrderUID] = order
	}
	return storage, nil
}

// Close - корректно закрывает соединение с базой данных.
func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) GetDb() *sqlx.DB {
	return s.db
}

// SaveOrder - сохранить заказ целиком или ничего.
func (s *Storage) SaveOrder(order *model.Order) error {
	// 1. Начинаем транзакцию.
	tx, err := s.db.Beginx()
	if err != nil {
		return fmt.Errorf("не удалось начать транзакцию: %w", err)
	}

	defer func() {
		if err := tx.Rollback(); !errors.Is(err, sql.ErrTxDone) {
			fmt.Printf("ошибка отката транзакции: %v", err)
		}
	}()

	_, err = tx.NamedExec(`
		INSERT INTO orders (
			order_uid, track_number, entry, locale, internal_signature, customer_id,
			delivery_service, shardkey, sm_id, date_created, oof_shard
		) VALUES (
			:order_uid, :track_number, :entry, :locale, :internal_signature, :customer_id,
			:delivery_service, :shardkey, :sm_id, :date_created, :oof_shard
		)
	`, order)
	if err != nil {
		return fmt.Errorf("не удалось сохранить заказ (order): %w", err)
	}

	dbDelivery := order.Delivery
	dbDelivery.OrderUID = order.OrderUID
	_, err = tx.NamedExec(`
		INSERT INTO deliveries (
			order_uid, name, phone, zip, city, address, region, email
		) VALUES (
			:order_uid, :name, :phone, :zip, :city, :address, :region, :email
		)
	`, &dbDelivery)
	if err != nil {
		return fmt.Errorf("не удалось сохранить доставку (delivery): %w", err)
	}

	dbPayment := order.Payment
	dbPayment.OrderUID = order.OrderUID
	_, err = tx.NamedExec(`
		INSERT INTO payments (
			order_uid, transaction, request_id, currency, provider, amount, payment_dt,
			bank, delivery_cost, goods_total, custom_fee
		) VALUES (
			:order_uid, :transaction, :request_id, :currency, :provider, :amount, :payment_dt,
			:bank, :delivery_cost, :goods_total, :custom_fee
		)
	`, &dbPayment)
	if err != nil {
		return fmt.Errorf("не удалось сохранить оплату (payment): %w", err)
	}

	for _, item := range order.Items {
		item.OrderUID = order.OrderUID
		_, err = tx.NamedExec(`
			INSERT INTO items (
				order_uid, chrt_id, track_number, price, rid, name, sale, size,
				total_price, nm_id, brand, status
			) VALUES (
				:order_uid, :chrt_id, :track_number, :price, :rid, :name, :sale, :size,
				:total_price, :nm_id, :brand, :status
			)
		`, &item)
		if err != nil {
			return fmt.Errorf("не удалось сохранить товар (item) с chrt_id %d: %w", item.ChrtID, err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("не удалось закоммитить транзакцию: %w", err)
	}

	return nil
}

// orderQueryResult - структура для получения ответа из SQL запроса.
type orderQueryResult struct {
	model.Order
	model.Delivery
	model.Payment
	ItemsJSON []byte `db:"items_json"`
}

// GetOrderByUID - получить заказ и связанные данные по uid.
func (s *Storage) GetOrderByUID(orderUID string) (*model.Order, error) {
	val, ok := s.cache[orderUID]
	if ok {
		return val, nil
	}

	sqlQuery := `
        SELECT
            o.*, d.*, p.*,
            COALESCE(
              (SELECT json_agg(i) FROM items i WHERE i.order_uid = o.order_uid),
              '[]'::json
            ) AS items_json
        FROM orders AS o
        LEFT JOIN deliveries AS d ON o.order_uid = d.order_uid
        LEFT JOIN payments AS p ON o.order_uid = p.order_uid
        WHERE o.order_uid = $1`

	var result orderQueryResult
	if err := s.db.Get(&result, sqlQuery, orderUID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("ошибка запроса к БД: %w", err)
	}

	var items []model.Item
	if err := json.Unmarshal(result.ItemsJSON, &items); err != nil {
		return nil, fmt.Errorf("ошибка распаковки JSON для товаров: %w", err)
	}

	order := &result.Order
	order.Delivery = result.Delivery
	order.Payment = result.Payment
	order.Items = items

	s.cache[orderUID] = order

	return order, nil
}

func (s *Storage) GetTopNewestOrders(quantity int) ([]*model.Order, error) {
	sqlQuery := `
        SELECT
            o.*, d.*, p.*,
            COALESCE(
              (SELECT json_agg(i) FROM items i WHERE i.order_uid = o.order_uid),
              '[]'::json
            ) AS items_json
        FROM orders AS o
        LEFT JOIN deliveries AS d ON o.order_uid = d.order_uid
        LEFT JOIN payments AS p ON o.order_uid = p.order_uid
        ORDER BY date_created DESC
        LIMIT $1`

	results := make([]orderQueryResult, quantity)
	if err := s.db.Select(&results, sqlQuery, quantity); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("ошибка запроса к БД: %w", err)
	}

	var items []model.Item
	orders := make([]*model.Order, len(results))

	for i := range results {
		result := results[i]
		if err := json.Unmarshal(result.ItemsJSON, &items); err != nil {
			return nil, fmt.Errorf("ошибка распаковки JSON для товаров: %w", err)
		}

		order := &result.Order
		order.Delivery = result.Delivery
		order.Payment = result.Payment
		order.Items = items
		orders[i] = order
	}
	return orders, nil
}
