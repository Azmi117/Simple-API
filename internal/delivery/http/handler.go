package http

import (
	"encoding/json"
	"net/http"

	"github.com/Azmi117/Simple-API/internal/models"
	"github.com/Azmi117/Simple-API/internal/usecase"
)

type ProductHandler struct {
	useCase *usecase.ProductUseCase
}

func NewProductHandler(u *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		useCase: u,
	}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products := h.useCase.GetAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.Product
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Input busuk cuy", 400)
		return
	}

	res, err := h.useCase.Create(input)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
