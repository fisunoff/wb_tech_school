package model

import (
	"testing"
)

func TestNewFakeOrder_GeneratesValidOrderWithoutError(t *testing.T) {
	order, err := NewFakeOrder()
	if err != nil {
		t.Error(err.Error())
	}
	if order == nil {
		t.Fatal("Order is nil")
	}
	if order.OrderUID == "" {
		t.Fatal("OrderUID пустой")
	}
	if order.TrackNumber == "" {
		t.Fatal("TrackNumber пустой")
	}
	if len(order.Items) == 0 {
		t.Fatal("Items пустой (должен быть хотя бы один элемент)")
	}
	if order.Items == nil {
		t.Fatal("Items не заполнились")
	}

	goodsTotal := 0
	for _, item := range order.Items {
		expectedTotalPrice := item.Price * (100 - item.Sale) / 100
		if item.TotalPrice != expectedTotalPrice {
			t.Errorf("Item.TotalPrice неверно рассчитан: ожидалось %d, получено %d", expectedTotalPrice, item.TotalPrice)
		}
		goodsTotal += item.TotalPrice
	}

	if order.Payment.GoodsTotal != goodsTotal {
		t.Errorf("Payment.GoodsTotal неверно: ожидалось %d, получено %d", goodsTotal, order.Payment.GoodsTotal)
	}

	expectedAmount := goodsTotal + order.Payment.DeliveryCost + order.Payment.CustomFee
	if order.Payment.Amount != expectedAmount {
		t.Errorf("Payment.Amount неверно: ожидалось %d, получено %d", expectedAmount, order.Payment.Amount)
	}
}
