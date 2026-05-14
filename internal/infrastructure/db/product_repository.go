package db

import (
	"context"
	"database/sql"

	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/jnates/crud_golang/internal/domain/ports"
	"github.com/jnates/crud_golang/internal/infrastructure/db/queries"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ports.ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, product *model.Product) error {
	_, err := r.db.ExecContext(ctx, queries.QueryCreateProduct,
		product.ID, product.Name, product.Description, product.Price, product.Stock, product.ImageURL, product.Category, product.SKU, product.IsActive, product.CreatedAt, product.UpdatedAt)
	return err
}

func (r *ProductRepository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	err := r.db.QueryRowContext(ctx, queries.QueryGetProductByID, id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.ImageURL, &product.Category, &product.SKU, &product.IsActive, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(ctx context.Context, product *model.Product) error {
	_, err := r.db.ExecContext(ctx, queries.QueryUpdateProduct,
		product.Name, product.Description, product.Price, product.Stock, product.ImageURL, product.Category, product.SKU, product.IsActive, product.UpdatedAt, product.ID)
	return err
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, queries.QueryDeleteProduct, id)
	return err
}

func (r *ProductRepository) List(ctx context.Context) ([]model.Product, error) {
	rows, err := r.db.QueryContext(ctx, queries.QueryListProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.ImageURL, &product.Category, &product.SKU, &product.IsActive, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
