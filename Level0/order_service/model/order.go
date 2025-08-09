package model

import (
	"encoding/json"
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid" db:"order_uid" fake:"{regex:[a-z0-9]{14,20}}"`
	TrackNumber       string    `json:"track_number" db:"track_number" fake:"{regex:[A-Z0-9]{14,20}}"`
	Entry             string    `json:"entry" db:"entry" fake:"{regex:[A-Z]{4}}"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale" db:"locale" fake:"{randomstring:[en,ru]}"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature" fake:"skip"`
	CustomerID        string    `json:"customer_id" db:"customer_id" fake:"{regex:[a-z0-9]{5,20}}"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service" fake:"{randomstring:[meest,cdek,boxberry,dpd,dhl,ups]}"`
	ShardKey          string    `json:"shardkey" db:"shardkey" fake:"{randomstring:[1,2,3,4,5,6,7,8,9,10]}"`
	SmID              int       `json:"sm_id" db:"sm_id" fake:"{number:1,100}"`
	DateCreated       time.Time `json:"date_created" db:"date_created" fake:"{date}"`
	OofShard          string    `json:"oof_shard" db:"oof_shard" fake:"{randomstring:[1,2,3,4,5,6,7,8,9,10]}"`
}

type Delivery struct {
	OrderUID string `db:"order_uid" fake:"skip"` // свяжем в постпроцессоре
	Name     string `json:"name" db:"name" fake:"{name}"`
	Phone    string `json:"phone" db:"phone" fake:"{phone}"`
	Zip      string `json:"zip" db:"zip" fake:"{zip}"`
	City     string `json:"city" db:"city" fake:"{city}"`
	Address  string `json:"address" db:"address" fake:"{street}"`
	Region   string `json:"region" db:"region" fake:"{state}"`
	Email    string `json:"email" db:"email" fake:"{email}"`
}

type Payment struct {
	OrderUID     string `db:"order_uid" fake:"skip"` // свяжем в постпроцессоре
	Transaction  string `json:"transaction" db:"transaction" fake:"{regex:[a-z0-9]{14,20}}"`
	RequestID    string `json:"request_id" db:"request_id" fake:"skip"`
	Currency     string `json:"currency" db:"currency" fake:"{randomstring:[USD,EUR,RUB]}"`
	Provider     string `json:"provider" db:"provider" fake:"{randomstring:[wbpay,yookassa,stripe,paypal]}"`
	Amount       int    `json:"amount" db:"amount" fake:"{number:100,500000}"`
	PaymentDT    int64  `json:"payment_dt" db:"payment_dt" fake:"{number:100000,1000000}"`
	Bank         string `json:"bank" db:"bank" fake:"{randomstring:[alpha,sber,tinkoff,vtb]}"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost" fake:"{number:0,3000}"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total" fake:"{number:100,100000}"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee" fake:"{number:0,200}"`
}

type Item struct {
	ID          int64  `db:"id" fake:"skip"`
	OrderUID    string `db:"order_uid" fake:"skip"`
	ChrtID      int64  `json:"chrt_id" db:"chrt_id" fake:"{number:1000000,10000000}"`
	TrackNumber string `json:"track_number" db:"track_number" fake:"{regex:[A-Z0-9]{14,20}}"`
	Price       int    `json:"price" db:"price" fake:"{number:100,5000}"`
	RID         string `json:"rid" db:"rid" fake:"{regex:[a-f0-9]{20}}"`
	Name        string `json:"name" db:"name" fake:"{randomstring:[Mascara,Lipstick,Eyeshadow,Concealer,Foundation]}"`
	Sale        int    `json:"sale" db:"sale" fake:"{number:0,70}"`
	Size        string `json:"size" db:"size" fake:"{randomstring:[XS,S,M,L,XL]}"`
	TotalPrice  int    `json:"total_price" db:"total_price" fake:"skip"` // постпроцессор
	NmID        int64  `json:"nm_id" db:"nm_id" fake:"{number:100000,99999999}"`
	Brand       string `json:"brand" db:"brand" fake:"{company}"`
	Status      int    `json:"status" db:"status" fake:"{number:100,300}"`
}

func SerializeOrder(orderRawData []byte) (Order, error) {
	var order Order
	err := json.Unmarshal(orderRawData, &order)
	return order, err
}
