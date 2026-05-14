package ports

import (
	"context"

	"github.com/jnates/crud_golang/internal/domain/model"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	GetByID(ctx context.Context, id string) (*model.Product, error)
	Update(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]model.Product, error)
}
