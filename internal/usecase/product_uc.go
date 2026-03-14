package usecase

import (
	"errors"
	"time"

	"github.com/Azmi117/Simple-API/internal/models"
	"github.com/Azmi117/Simple-API/internal/repository"
)

type ProductUseCase struct {
	repo *repository.ProductRepository
}

func NewProductUseCase(r *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repo: r,
	}
}

func (u *ProductUseCase) GetAll() []models.Product {
	return u.repo.FindAll()
}

func (u *ProductUseCase) Create(p models.Product) (models.Product, error) {
	if p.Name == "" {
		return models.Product{}, errors.New("nama produk tidak boleh kosong cuy")
	}
	if p.Price <= 0 {
		return models.Product{}, errors.New("harga harus lebih dari 0")
	}

	// 2. Set Timestamps (Biar real world)
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	// 3. Set ID (Karena kita gak pake DB Auto Increment, kita manual dulu)
	// Untuk latihan ini, lu bisa isi ID sembarang atau pake logic len(slice)+1
	p.ID = len(u.repo.FindAll()) + 1

	return u.repo.Create(p), nil
}
