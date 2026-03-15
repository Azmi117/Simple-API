package http

import "net/http"

func MapRoutes(mux *http.ServeMux, handler *ProductHandler) {
	mux.HandleFunc("GET /products", handler.GetAll)
	mux.HandleFunc("POST /products", handler.Create)
	mux.HandleFunc("PATCH /products/{id}", handler.Update)
	mux.HandleFunc("DELETE /products/{id}", handler.Delete)

	// Nanti tinggal nambah di sini, main.go tetep bersih
	// mux.HandleFunc("PATCH /products/{id}", handler.Update)
	// mux.HandleFunc("DELETE /products/{id}", handler.Delete)
}
