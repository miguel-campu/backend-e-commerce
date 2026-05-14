package application

import (
	"context"
	"testing"

	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) List(ctx context.Context) ([]model.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductRepository) GetByID(ctx context.Context, id string) (*model.Product, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductRepository) Create(ctx context.Context, p *model.Product) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *MockProductRepository) Update(ctx context.Context, p *model.Product) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProductService(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()

	t.Run("ListProducts", func(t *testing.T) {
		products := []model.Product{{Name: "P1"}}
		mockRepo.On("List", ctx).Return(products, nil)
		res, err := service.ListProducts(ctx)
		assert.NoError(t, err)
		assert.Equal(t, products, res)
	})

	t.Run("CreateProduct", func(t *testing.T) {
		p := &model.Product{Name: "New"}
		mockRepo.On("Create", ctx, mock.AnythingOfType("*model.Product")).Return(nil)
		err := service.CreateProduct(ctx, p)
		assert.NoError(t, err)
	})

	t.Run("GetProduct", func(t *testing.T) {
		p := &model.Product{Name: "P1"}
		mockRepo.On("GetByID", ctx, "1").Return(p, nil)
		res, err := service.GetProduct(ctx, "1")
		assert.NoError(t, err)
		assert.Equal(t, p, res)
	})

	t.Run("UpdateProduct", func(t *testing.T) {
		p := &model.Product{Name: "Updated"}
		mockRepo.On("Update", ctx, p).Return(nil)
		err := service.UpdateProduct(ctx, p)
		assert.NoError(t, err)
	})

	t.Run("DeleteProduct", func(t *testing.T) {
		mockRepo.On("Delete", ctx, "1").Return(nil)
		err := service.DeleteProduct(ctx, "1")
		assert.NoError(t, err)
	})
}
