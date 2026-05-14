package application

import (
	"context"
	"errors"
	"testing"

	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/jnates/crud_golang/internal/domain/ports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPaymentProvider struct {
	mock.Mock
}

func (m *MockPaymentProvider) GeneratePaymentLink(order *model.Order, userEmail string) (*ports.PaymentResponse, error) {
	args := m.Called(order, userEmail)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ports.PaymentResponse), args.Error(1)
}

func TestPaymentService(t *testing.T) {
	mockRepo := new(MockOrdersRepository)
	mockGateway := new(MockPaymentProvider)
	service := NewPaymentService(mockRepo, mockGateway)
	ctx := context.Background()

	t.Run("ProcessPayment Success", func(t *testing.T) {
		order := &model.Order{OrderNumber: "ORD-123"}
		email := "test@example.com"
		response := &ports.PaymentResponse{
			CheckoutURL:   "http://payment.com/pay",
			TransactionID: "TX-1",
		}

		mockRepo.On("GetOrderByID", ctx, "order-id").Return(order, nil).Once()
		mockGateway.On("GeneratePaymentLink", order, email).Return(response, nil).Once()

		res, err := service.ProcessPayment(ctx, "order-id", email)

		assert.NoError(t, err)
		assert.Equal(t, response, res)
		mockRepo.AssertExpectations(t)
		mockGateway.AssertExpectations(t)
	})

	t.Run("ProcessPayment Order Not Found", func(t *testing.T) {
		mockRepo.On("GetOrderByID", ctx, "invalid-id").Return(nil, errors.New("order not found")).Once()

		res, err := service.ProcessPayment(ctx, "invalid-id", "test@example.com")

		assert.Error(t, err)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("ProcessPayment Gateway Error", func(t *testing.T) {
		order := &model.Order{OrderNumber: "ORD-456"}
		email := "error@example.com"

		mockRepo.On("GetOrderByID", ctx, "order-id").Return(order, nil).Once()
		mockGateway.On("GeneratePaymentLink", order, email).Return(nil, errors.New("gateway error")).Once()

		res, err := service.ProcessPayment(ctx, "order-id", email)

		assert.Error(t, err)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
		mockGateway.AssertExpectations(t)
	})
}
