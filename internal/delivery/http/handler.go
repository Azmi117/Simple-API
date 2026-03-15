package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

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

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID harus angka, Cuy!", http.StatusBadRequest)
		return
	}

	var input models.Product
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Format JSON lu rusak", http.StatusBadRequest)
		return
	}

	res, err := h.useCase.Update(id, input)
	if err != nil {
		// CEK JENIS ERROR-NYA
		if errors.Is(err, usecase.ErrProductNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound) // 404
			return
		}

		// Kalau error lain, anggap bad request
		http.Error(w, err.Error(), http.StatusBadRequest) // 400
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Tambahan: Handler Delete
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	err := h.useCase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
