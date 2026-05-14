package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID            uuid.UUID   `json:"id"`
	OrderNumber   string      `json:"orderNumber"`
	UserID        uuid.UUID   `json:"customerId"`
	CustomerName  string      `json:"customerName"`
	CustomerEmail string      `json:"customerEmail"`
	TotalAmount   float64     `json:"total"`
	Status        string      `json:"status"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	Items         []OrderItem `json:"items,omitempty"`
}

type OrderItem struct {
	ID          uuid.UUID `json:"id"`
	OrderID     uuid.UUID `json:"orderId"`
	ProductID   uuid.UUID `json:"productId"`
	ProductName string    `json:"productName"`
	Quantity    int       `json:"quantity"`
	PriceAtTime float64   `json:"unitPrice"`
}
