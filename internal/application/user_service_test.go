package application

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock of ports.UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdatePassword(ctx context.Context, id string, passwordHash string) error {
	args := m.Called(ctx, id, passwordHash)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) List(ctx context.Context) ([]model.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.User), args.Error(1)
}

func TestUserService_Login(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		email := "test@example.com"
		password := "password123"
		// In a real scenario, you'd need the actual hash from the tool kit
		// For simplicity, we mock the behavior of GetByEmail
		user := &model.User{
			ID:           uuid.New(),
			Email:        email,
			PasswordHash: "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f", // Correct hash for password123
			FullName:     "Test User",
			Role:         "user",
		}

		mockRepo.On("GetByEmail", ctx, email).Return(user, nil)

		res, err := service.Login(ctx, model.LoginRequest{Email: email, Password: password})

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, email, res.User.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		email := "nonexistent@example.com"
		mockRepo.On("GetByEmail", ctx, email).Return(nil, assert.AnError)

		res, err := service.Login(ctx, model.LoginRequest{Email: email, Password: "any"})

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Contains(t, err.Error(), "inválidos")
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		email := "test@example.com"
		user := &model.User{
			Email:        email,
			PasswordHash: "wronghash",
		}
		mockRepo.On("GetByEmail", ctx, email).Return(user, nil)

		res, err := service.Login(ctx, model.LoginRequest{Email: email, Password: "wrongpassword"})

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Contains(t, err.Error(), "inválidos")
	})
}

func TestUserService_Register(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		req := model.RegisterRequest{
			Email:    "new@example.com",
			Password: "password123",
			FullName: "New User",
		}
		mockRepo.On("GetByEmail", ctx, req.Email).Return(nil, nil)
		mockRepo.On("Create", ctx, mock.AnythingOfType("*model.User")).Return(nil)

		res, err := service.Register(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, req.Email, res.User.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UserAlreadyExists", func(t *testing.T) {
		req := model.RegisterRequest{Email: "exists@example.com"}
		mockRepo.On("GetByEmail", ctx, req.Email).Return(&model.User{}, nil)

		res, err := service.Register(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "user already exists", err.Error())
	})
}

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		req := model.CreateUserRequest{
			Email:    "admin@example.com",
			Password: "password123",
			FullName: "Admin User",
			Role:     "admin",
		}
		mockRepo.On("GetByEmail", ctx, req.Email).Return(nil, nil)
		mockRepo.On("Create", ctx, mock.AnythingOfType("*model.User")).Return(nil)

		res, err := service.CreateUser(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "admin", res.Role)
	})
}

func TestUserService_UpdatePassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		id := "123"
		newPass := "newpassword"
		mockRepo.On("UpdatePassword", ctx, id, mock.AnythingOfType("string")).Return(nil)

		err := service.UpdatePassword(ctx, id, newPass)

		assert.NoError(t, err)
	})
}

func TestUserService_CRUD(t *testing.T) {
	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)
	ctx := context.Background()

	t.Run("GetUser", func(t *testing.T) {
		id := uuid.New().String()
		user := &model.User{ID: uuid.MustParse(id)}
		mockRepo.On("GetByID", ctx, id).Return(user, nil)
		res, err := service.GetUser(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, user, res)
	})

	t.Run("ListUsers", func(t *testing.T) {
		users := []model.User{{FullName: "User 1"}}
		mockRepo.On("List", ctx).Return(users, nil)
		res, err := service.ListUsers(ctx)
		assert.NoError(t, err)
		assert.Equal(t, users, res)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		user := &model.User{FullName: "Updated"}
		mockRepo.On("Update", ctx, user).Return(nil)
		err := service.UpdateUser(ctx, user)
		assert.NoError(t, err)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		id := "123"
		mockRepo.On("Delete", ctx, id).Return(nil)
		err := service.DeleteUser(ctx, id)
		assert.NoError(t, err)
	})
}
