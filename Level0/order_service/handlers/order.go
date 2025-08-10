package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"order_service/storage"

	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
	DB *storage.Storage
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
	if orderUID == "" {
		http.Error(w, "Пустой uid", http.StatusBadRequest)
		return
	}

	order, err := h.DB.GetOrderByUID(orderUID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Заказ не найден", http.StatusNotFound)
			return
		}
		log.Printf("Ошибка при попытке получить заказ %s: %v", orderUID, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(order); err != nil {
		log.Printf("Ошибка при записи ответа: %v", err)
	}
}
