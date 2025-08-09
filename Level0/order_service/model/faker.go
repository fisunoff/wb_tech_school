package model

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func NewFakeOrder() (Order, error) {
	// По-хорошему seed делать один раз в init/main, но оставлю как было
	err := gofakeit.Seed(time.Now().UnixNano())
	if err != nil {
		return Order{}, err
	}

	n := gofakeit.Number(1, 50)
	o := Order{
		Items: make([]Item, n),
	}

	err = gofakeit.Struct(&o)
	if err != nil {
		return Order{}, err
	}
	postProcessOrder(&o)
	return o, nil
}

func postProcessOrder(o *Order) {
	// Принимаем, что UID совпадают
	o.Payment.OrderUID = o.OrderUID
	o.Payment.Transaction = o.OrderUID
	o.Delivery.OrderUID = o.OrderUID

	// Пересчёт товаров и итогов
	goodsTotal := 0
	for i := range o.Items {
		it := &o.Items[i]
		it.OrderUID = o.OrderUID
		it.TrackNumber = o.TrackNumber

		it.TotalPrice = it.Price * (100 - it.Sale) / 100
		goodsTotal += it.TotalPrice
	}

	o.Payment.GoodsTotal = goodsTotal
	o.Payment.Amount = o.Payment.GoodsTotal + o.Payment.DeliveryCost + o.Payment.CustomFee
}
