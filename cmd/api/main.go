package main

import (
	"fmt"
	"net/http"

	delivery "github.com/Azmi117/Simple-API/internal/delivery/http"
	"github.com/Azmi117/Simple-API/internal/repository"
	"github.com/Azmi117/Simple-API/internal/usecase"
)

func main() {
	// 1. Setup Layer
	repo := repository.NewProductRepository(repository.DB)
	uc := usecase.NewProductUseCase(repo)
	handler := delivery.NewProductHandler(uc)

	// 2. Setup Router
	mux := http.NewServeMux()

	// 3. Panggil Mapping Routes dari file sebelah
	delivery.MapRoutes(mux, handler)

	// 4. Run Server
	port := ":8080"
	fmt.Printf("Server lari kenceng di port %s...\n", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Waduh, servernya mogok:", err)
	}
}
