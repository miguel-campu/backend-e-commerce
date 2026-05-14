package application

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/jnates/crud_golang/internal/domain/ports"
)

type OrderService struct {
	repo ports.OrdersRepository
}

func NewOrderService(repo ports.OrdersRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) List(ctx context.Context, userID string, isAdmin bool, status string) ([]model.Order, error) {
	if isAdmin {
		return s.repo.ListOrders(ctx, status)
	}
	return s.repo.ListOrdersByUserID(ctx, userID, status)
}

func (s *OrderService) ListOrderItems(ctx context.Context) ([]model.OrderItem, error) {
	return s.repo.ListOrderItems(ctx)
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (*model.Order, error) {
	return s.repo.GetOrderByID(ctx, id)
}

func (s *OrderService) Create(ctx context.Context, order *model.Order) error {
	order.ID = uuid.New()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	if order.Status == "" {
		order.Status = "pending"
	}

	// Fallback: If total is 0, calculate it from items
	if order.TotalAmount <= 0 {
		var calcTotal float64
		for _, item := range order.Items {
			calcTotal += float64(item.Quantity) * item.PriceAtTime
		}
		order.TotalAmount = calcTotal
	}

	return s.repo.CreateOrder(ctx, order)
}

func (s *OrderService) Update(ctx context.Context, order *model.Order) error {
	order.UpdatedAt = time.Now()
	return s.repo.UpdateOrder(ctx, order)
}

func (s *OrderService) Delete(ctx context.Context, id string) error {
	return s.repo.DeleteOrder(ctx, id)
}

//servivio de prueba del pull request
