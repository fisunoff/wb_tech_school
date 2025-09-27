package usecase

import (
	"errors"
	"testing"

	"order_service/internal/model"
	"order_service/internal/repository"
)

type mockOrderRepo struct {
	OnGetByUID func(uid string) (*model.Order, error)
}

func (m *mockOrderRepo) FindByUID(uid string) (*model.Order, error) {
	if m.OnGetByUID != nil {
		return m.OnGetByUID(uid)
	}
	return nil, errors.New("not implemented")
}

func (m *mockOrderRepo) Save(order *model.Order) error {
	return errors.New("not implemented")
}

func (m *mockOrderRepo) FindTopNewest(quantity int) ([]*model.Order, error) {
	return nil, errors.New("not implemented")
}

func TestOrderUseCase_GetOrderByUID_EmptyUID(t *testing.T) {
	repo := &mockOrderRepo{}
	uc := NewOrderUseCase(repo)

	_, err := uc.GetOrderByUID("")
	if !errors.Is(err, ErrEmptyUid) {
		t.Errorf("Ожидалась ошибка ErrEmptyUid, получена: %v", err)
	}
}

func TestOrderUseCase_GetOrderByUID_NotFound(t *testing.T) {
	repo := &mockOrderRepo{
		OnGetByUID: func(uid string) (*model.Order, error) {
			return nil, repository.ErrNotFound
		},
	}
	uc := NewOrderUseCase(repo)

	_, err := uc.GetOrderByUID("nonexistent")
	if !errors.Is(err, ErrOrderNotFound) {
		t.Errorf("Ожидалась ошибка ErrOrderNotFound, получена: %v", err)
	}
}

func TestOrderUseCase_GetOrderByUID_Success(t *testing.T) {
	expectedOrder := &model.Order{OrderUID: "test123"}
	repo := &mockOrderRepo{
		OnGetByUID: func(uid string) (*model.Order, error) {
			return expectedOrder, nil
		},
	}
	uc := NewOrderUseCase(repo)

	order, err := uc.GetOrderByUID("test123")
	if err != nil {
		t.Fatal(err)
	}
	if order != expectedOrder {
		t.Error("Возвращён неверный заказ")
	}
}
