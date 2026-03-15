package usecase

import (
	"errors"
	"time"

	"github.com/Azmi117/Simple-API/internal/models"
	"github.com/Azmi117/Simple-API/internal/repository"
)

var (
	ErrProductNotFound = errors.New("produknya ga ada, Mi!")
	ErrInvalidInput    = errors.New("input lu ada yang salah nih")
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

func (u *ProductUseCase) Update(id int, input models.Product) (models.Product, error) {
	existing, found := u.repo.FindByID(id)
	if !found {
		return models.Product{}, ErrProductNotFound
	}

	// PATCH logic: cuma ganti kalau ada isinya
	if input.Name != "" {
		existing.Name = input.Name
	} else {
		return models.Product{}, ErrInvalidInput
	}

	if input.Price > 0 {
		existing.Price = input.Price
	} else {
		return models.Product{}, ErrInvalidInput
	}

	existing.UpdatedAt = time.Now()
	u.repo.Update(existing)
	return existing, nil
}

// Tambahan: Logic Delete
func (u *ProductUseCase) Delete(id int) error {
	success := u.repo.Delete(id)
	if !success {
		return errors.New("gagal hapus, ID ga ketemu")
	}
	return nil
}
