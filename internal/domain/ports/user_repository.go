package ports

import (
	"context"

	"github.com/jnates/crud_golang/internal/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	UpdatePassword(ctx context.Context, id string, passwordHash string) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]model.User, error)
}
