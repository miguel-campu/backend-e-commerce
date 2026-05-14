package db

import (
	"context"
	"database/sql"

	"time"

	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/jnates/crud_golang/internal/domain/ports"
	"github.com/jnates/crud_golang/internal/infrastructure/db/queries"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, queries.QueryCreateUser,
		user.ID, user.Email, user.PasswordHash, user.Role, user.FullName, user.AvatarURL, user.GoogleID, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRowContext(ctx, queries.QueryGetUserByID, id).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Role, &user.FullName, &user.AvatarURL, &user.GoogleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRowContext(ctx, queries.QueryGetUserByEmail, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Role, &user.FullName, &user.AvatarURL, &user.GoogleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx, queries.QueryUpdateUser,
		user.Email, user.Role, user.FullName, user.AvatarURL, user.GoogleID, user.UpdatedAt, user.ID)
	return err
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id string, passwordHash string) error {
	_, err := r.db.ExecContext(ctx, queries.QueryUpdateUserPassword, passwordHash, time.Now(), id)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, queries.QueryDeleteUser, id)
	return err
}

func (r *UserRepository) List(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.QueryContext(ctx, queries.QueryListUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role, &user.FullName, &user.AvatarURL, &user.GoogleID, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
