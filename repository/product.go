package repository

import (
	"testgo/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	ListProductsPaginated(page, limit int, search string) ([]models.Product, int, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
}

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) *productRepository {
	return &productRepository{
		DB: DB,
	}
}


func (r *productRepository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error

	return products, err
}
func (r *productRepository) ListProductsPaginated(page, limit int, search string) ([]models.Product, int, error) {
	var products []models.Product
	var totalCount int64

	query := r.DB.Model(&models.Product{})
	if search != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	err = query.Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, int(totalCount), nil
}

func (r *productRepository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, ID).Error

	return product, err
}

func (r *productRepository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.DB.Create(&product).Error

	return product, err
}

func (r *productRepository) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.DB.Save(&product).Error

	return product, err
}

func (r *productRepository) DeleteProduct(product models.Product) (models.Product, error) {
	err := r.DB.Delete(&product).Error

	return product, err
}
