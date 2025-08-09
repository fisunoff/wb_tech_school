package model

import (
	"testing"
)

// TestFaker - остановимся на том, что хотя бы какие-то данные есть и нет ошибок
func TestFaker(t *testing.T) {
	order, err := NewFakeOrder()
	if err != nil {
		t.Error(err.Error())
	}
	if order.Items == nil {
		t.Error("Items не заполнились")
	}
}
