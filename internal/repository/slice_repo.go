package repository

import (
	"time"

	"github.com/Azmi117/Simple-API/internal/models"
)

type MemDb struct {
	Product []models.Product
}

var DB = &MemDb{
	Product: []models.Product{},
}

type ProductRepository struct {
	db *MemDb
}

func NewProductRepository(database *MemDb) *ProductRepository {
	return &ProductRepository{
		db: database,
	}
}

// Pastikan huruf F-nya KAPITAL supaya bisa dipanggil dari folder usecase
func (r *ProductRepository) FindAll() []models.Product {
	var activeProducts []models.Product

	// Kita looping isi database-nya
	for _, p := range r.db.Product {
		// Cek logic Soft Delete: cuma ambil yang DeletedAt-nya nil
		if p.DeletedAt == nil {
			activeProducts = append(activeProducts, p)
		}
	}

	return activeProducts
}

func (r *ProductRepository) Create(p models.Product) models.Product {
	r.db.Product = append(r.db.Product, p)
	return p
}

func (r *ProductRepository) FindByID(id int) (models.Product, bool) {
	for _, p := range r.db.Product {
		if p.ID == id && p.DeletedAt == nil {
			return p, true
		}
	}
	return models.Product{}, false
}

// Tambahan: Update data di dalam slice
func (r *ProductRepository) Update(p models.Product) {
	for i, v := range r.db.Product {
		if v.ID == p.ID {
			r.db.Product[i] = p
			break
		}
	}
}

// Tambahan: Soft Delete (Cuma kasih cap waktu)
func (r *ProductRepository) Delete(id int) bool {
	for i, p := range r.db.Product {
		if p.ID == id && p.DeletedAt == nil {
			now := time.Now()
			r.db.Product[i].DeletedAt = &now
			return true
		}
	}
	return false
}
