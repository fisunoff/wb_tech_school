package model

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
)

type Order struct {
	OrderUID          string    `json:"order_uid" db:"order_uid" fake:"{regex:[a-z0-9]{14,20}}" validate:"required"`
	TrackNumber       string    `json:"track_number" db:"track_number" fake:"{regex:[A-Z0-9]{14,20}}" validate:"required"`
	Entry             string    `json:"entry" db:"entry" fake:"{regex:[A-Z]{4}}" validate:"required"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" validate:"required,min=1,dive"`
	Locale            string    `json:"locale" db:"locale" fake:"{randomstring:[en,ru]}" validate:"required,oneof=en ru"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature" fake:"skip"`
	CustomerID        string    `json:"customer_id" db:"customer_id" fake:"{regex:[a-z0-9]{5,20}}" validate:"required"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service" fake:"{randomstring:[meest,cdek,boxberry,dpd,dhl,ups]}" validate:"required"`
	ShardKey          string    `json:"shardkey" db:"shardkey" fake:"{randomstring:[1,2,3,4,5,6,7,8,9,10]}" validate:"required"`
	SmID              int       `json:"sm_id" db:"sm_id" fake:"{number:1,100}" validate:"required,gt=0"`
	DateCreated       time.Time `json:"date_created" db:"date_created" fake:"{date}" validate:"required"`
	OofShard          string    `json:"oof_shard" db:"oof_shard" fake:"{randomstring:[1,2,3,4,5,6,7,8,9,10]}" validate:"required"`
}

type Delivery struct {
	OrderUID string `db:"order_uid" fake:"skip"`
	Name     string `json:"name" db:"name" fake:"{name}" validate:"required"`
	Phone    string `json:"phone" db:"phone" fake:"{phone}" validate:"required"`
	Zip      string `json:"zip" db:"zip" fake:"{zip}" validate:"required"`
	City     string `json:"city" db:"city" fake:"{city}" validate:"required"`
	Address  string `json:"address" db:"address" fake:"{street}" validate:"required"`
	Region   string `json:"region" db:"region" fake:"{state}" validate:"required"`
	Email    string `json:"email" db:"email" fake:"{email}" validate:"required,email"`
}

type Payment struct {
	OrderUID     string `db:"order_uid" fake:"skip"`
	Transaction  string `json:"transaction" db:"transaction" fake:"{regex:[a-z0-9]{14,20}}" validate:"required"`
	RequestID    string `json:"request_id" db:"request_id" fake:"skip"`
	Currency     string `json:"currency" db:"currency" fake:"{randomstring:[USD,EUR,RUB]}" validate:"required,oneof=USD EUR RUB"`
	Provider     string `json:"provider" db:"provider" fake:"{randomstring:[wbpay,yookassa,stripe,paypal]}" validate:"required"`
	Amount       int    `json:"amount" db:"amount" fake:"{number:100,500000}" validate:"gt=0"`
	PaymentDT    int64  `json:"payment_dt" db:"payment_dt" fake:"{number:100000,1000000}" validate:"required,gt=0"`
	Bank         string `json:"bank" db:"bank" fake:"{randomstring:[alpha,sber,tinkoff,vtb]}" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost" fake:"{number:0,3000}" validate:"gte=0"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total" fake:"{number:100,100000}" validate:"gt=0"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee" fake:"{number:0,200}" validate:"gte=0"`
}

type Item struct {
	ID          int64  `db:"id" fake:"skip"`
	OrderUID    string `db:"order_uid" fake:"skip"`
	ChrtID      int64  `json:"chrt_id" db:"chrt_id" fake:"{number:1000000,10000000}" validate:"required,gt=0"`
	TrackNumber string `json:"track_number" db:"track_number" fake:"{regex:[A-Z0-9]{14,20}}" validate:"required"`
	Price       int    `json:"price" db:"price" fake:"{number:100,5000}" validate:"gt=0"`
	RID         string `json:"rid" db:"rid" fake:"{regex:[a-f0-9]{20}}" validate:"required"`
	Name        string `json:"name" db:"name" fake:"{randomstring:[Mascara,Lipstick,Eyeshadow,Concealer,Foundation]}" validate:"required"`
	Sale        int    `json:"sale" db:"sale" fake:"{number:0,70}" validate:"gte=0,lte=100"`
	Size        string `json:"size" db:"size" fake:"{randomstring:[XS,S,M,L,XL]}"`
	TotalPrice  int    `json:"total_price" db:"total_price" fake:"skip" validate:"gte=0"`
	NmID        int64  `json:"nm_id" db:"nm_id" fake:"{number:100000,99999999}" validate:"required,gt=0"`
	Brand       string `json:"brand" db:"brand" fake:"{company}" validate:"required"`
	Status      int    `json:"status" db:"status" fake:"{number:100,300}" validate:"required,gt=0"`
}

// SerializeOrder десериализует JSON и проверяет его валидность
func SerializeOrder(orderRawData []byte) (Order, error) {
	var order Order
	err := json.Unmarshal(orderRawData, &order)
	if err != nil {
		return order, err
	}

	validate := validator.New()
	err = validate.Struct(order)
	if err != nil {
		return order, err
	}

	return order, nil
}
