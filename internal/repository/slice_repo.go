package repository

import "github.com/Azmi117/Simple-API/internal/models"

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
