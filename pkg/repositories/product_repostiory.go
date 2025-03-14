package repositories

import (
	"database/sql"
	"ecomm-crud/pkg/config/db"
	"ecomm-crud/pkg/models"
	"errors"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository() *ProductRepository {
	db := db.CreateDBConnection()
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product *models.Product) (*models.Product, error) {

	query := "INSERT INTO products (name, image_url, stock, record_status) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRow(query, product.Name, product.ImageUrl, product.Stock, product.RecordStatus).Scan(&product.ID)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) GetProductByID(id int) (*models.Product, error) {

	query := "SELECT id, name, image_url, stock, record_status FROM products where id = $1"
	row := r.db.QueryRow(query, id)
	if row != nil {
		fmt.Println(row.Err().Error())
		var product models.Product
		err := row.Scan(&product.ID, &product.Name, &product.ImageUrl, &product.Stock, &product.RecordStatus)
		if err != nil {
			return nil, err
		}
		return &product, nil
	}
	return nil, errors.New("No product found with id")

}
