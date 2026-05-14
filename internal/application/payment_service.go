package application

import (
	"context"

	"github.com/jnates/crud_golang/internal/domain/ports"
)

type PaymentService struct {
	orderRepo ports.OrdersRepository
	gateway   ports.PaymentProvider
}

func NewPaymentService(repo ports.OrdersRepository, gw ports.PaymentProvider) *PaymentService {
	return &PaymentService{orderRepo: repo, gateway: gw}
}

func (s *PaymentService) ProcessPayment(ctx context.Context, orderID string, email string) (*ports.PaymentResponse, error) {
	order, err := s.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	return s.gateway.GeneratePaymentLink(order, email)
}
