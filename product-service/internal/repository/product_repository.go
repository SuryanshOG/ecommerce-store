package repository

import (
	"database/sql"
	"product-service/internal/model"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) CreateProduct(product model.Product) error {
	query := "INSERT INTO products (name, price) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, product.Name, product.Price)
	return err
}

func (r *ProductRepository) GetAllProducts() ([]model.Product, error) {
	query := "SELECT id, name, price FROM products"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
