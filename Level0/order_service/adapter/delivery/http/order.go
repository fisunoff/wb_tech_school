package http

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"order_service/internal/usecase"
)

type OrderHandler struct {
	useCase *usecase.OrderUseCase
}

func NewOrderHandler(useCase *usecase.OrderUseCase) *OrderHandler {
	return &OrderHandler{useCase}
}

// GetByUID - обработчик для GET /order/{order_uid}.
//
// @Summary      Получить заказ по UID
// @Description  Выдает полную информацию о заказе по UID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order_uid   path      string  true  "UID Заказа"
// @Success      200  {object}  model.Order
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /order/{order_uid} [get]
func (h *OrderHandler) GetByUID(w http.ResponseWriter, r *http.Request) {
	orderUID := chi.URLParam(r, "order_uid")

	order, err := h.useCase.GetOrderByUID(orderUID)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrEmptyUid):
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case errors.Is(err, usecase.ErrOrderNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		log.Printf("Ошибка при записи ответа: %v", err)
	}
}
