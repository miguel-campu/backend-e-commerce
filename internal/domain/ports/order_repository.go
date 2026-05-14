package ports

import (
	"context"

	"github.com/jnates/crud_golang/internal/domain/model"
)

type OrdersRepository interface {
	ListOrders(ctx context.Context, status string) ([]model.Order, error)
	ListOrdersByUserID(ctx context.Context, userID string, status string) ([]model.Order, error)
	ListOrderItems(ctx context.Context) ([]model.OrderItem, error)
	GetOrderByID(ctx context.Context, id string) (*model.Order, error)
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context, order *model.Order) error
	DeleteOrder(ctx context.Context, id string) error
}
