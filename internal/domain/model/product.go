package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"required,min=3"`
	Description string    `json:"description" validate:"required"`
	Price       float64   `json:"price" validate:"required,gt=0"`
	Stock       int       `json:"stock" validate:"required,gte=0"`
	ImageURL    string    `json:"imageUrl"`
	Category    string    `json:"category" validate:"required"`
	SKU         string    `json:"sku" validate:"required"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
