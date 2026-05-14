package application

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrdersRepository struct {
	mock.Mock
}

func (m *MockOrdersRepository) ListOrders(ctx context.Context, status string) ([]model.Order, error) {
	args := m.Called(ctx, status)
	return args.Get(0).([]model.Order), args.Error(1)
}

func (m *MockOrdersRepository) ListOrdersByUserID(ctx context.Context, userID string, status string) ([]model.Order, error) {
	args := m.Called(ctx, userID, status)
	return args.Get(0).([]model.Order), args.Error(1)
}

func (m *MockOrdersRepository) ListOrderItems(ctx context.Context) ([]model.OrderItem, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.OrderItem), args.Error(1)
}

func (m *MockOrdersRepository) GetOrderByID(ctx context.Context, id string) (*model.Order, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Order), args.Error(1)
}

func (m *MockOrdersRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *MockOrdersRepository) UpdateOrder(ctx context.Context, order *model.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *MockOrdersRepository) DeleteOrder(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestOrderService(t *testing.T) {
	mockRepo := new(MockOrdersRepository)
	service := NewOrderService(mockRepo)
	ctx := context.Background()

	t.Run("List Orders as Admin", func(t *testing.T) {
		orders := []model.Order{{OrderNumber: "ORD-1"}}
		mockRepo.On("ListOrders", ctx, "pending").Return(orders, nil).Once()

		res, err := service.List(ctx, "user-1", true, "pending")

		assert.NoError(t, err)
		assert.Equal(t, orders, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("List Orders as User", func(t *testing.T) {
		orders := []model.Order{{OrderNumber: "ORD-2"}}
		mockRepo.On("ListOrdersByUserID", ctx, "user-2", "completed").Return(orders, nil).Once()

		res, err := service.List(ctx, "user-2", false, "completed")

		assert.NoError(t, err)
		assert.Equal(t, orders, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("ListOrderItems", func(t *testing.T) {
		items := []model.OrderItem{{ProductName: "Item 1"}}
		mockRepo.On("ListOrderItems", ctx).Return(items, nil).Once()

		res, err := service.ListOrderItems(ctx)

		assert.NoError(t, err)
		assert.Equal(t, items, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetOrder Success", func(t *testing.T) {
		order := &model.Order{OrderNumber: "ORD-3"}
		mockRepo.On("GetOrderByID", ctx, "some-id").Return(order, nil).Once()

		res, err := service.GetOrder(ctx, "some-id")

		assert.NoError(t, err)
		assert.Equal(t, order, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetOrder Error", func(t *testing.T) {
		mockRepo.On("GetOrderByID", ctx, "fail-id").Return(nil, errors.New("not found")).Once()

		res, err := service.GetOrder(ctx, "fail-id")

		assert.Error(t, err)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create Order with Total", func(t *testing.T) {
		order := &model.Order{
			UserID:      uuid.New(),
			TotalAmount: 100.0,
		}
		mockRepo.On("CreateOrder", ctx, mock.AnythingOfType("*model.Order")).Return(nil).Once()

		err := service.Create(ctx, order)

		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, order.ID)
		assert.Equal(t, "pending", order.Status)
		assert.Equal(t, 100.0, order.TotalAmount)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create Order calculates Total", func(t *testing.T) {
		order := &model.Order{
			UserID: uuid.New(),
			Items: []model.OrderItem{
				{Quantity: 2, PriceAtTime: 10.0},
				{Quantity: 1, PriceAtTime: 5.0},
			},
		}
		mockRepo.On("CreateOrder", ctx, mock.AnythingOfType("*model.Order")).Return(nil).Once()

		err := service.Create(ctx, order)

		assert.NoError(t, err)
		assert.Equal(t, 25.0, order.TotalAmount)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create Order with predefined Status", func(t *testing.T) {
		order := &model.Order{
			UserID: uuid.New(),
			Status: "shipped",
		}
		mockRepo.On("CreateOrder", ctx, mock.AnythingOfType("*model.Order")).Return(nil).Once()

		err := service.Create(ctx, order)

		assert.NoError(t, err)
		assert.Equal(t, "shipped", order.Status)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create Order with predefined Total", func(t *testing.T) {
		order := &model.Order{
			UserID:      uuid.New(),
			TotalAmount: 500.0,
		}
		mockRepo.On("CreateOrder", ctx, mock.AnythingOfType("*model.Order")).Return(nil).Once()

		err := service.Create(ctx, order)

		assert.NoError(t, err)
		assert.Equal(t, 500.0, order.TotalAmount)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Update Order", func(t *testing.T) {
		order := &model.Order{OrderNumber: "ORD-UPDATE"}
		mockRepo.On("UpdateOrder", ctx, order).Return(nil).Once()

		err := service.Update(ctx, order)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Delete Order", func(t *testing.T) {
		mockRepo.On("DeleteOrder", ctx, "del-id").Return(nil).Once()

		err := service.Delete(ctx, "del-id")

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
