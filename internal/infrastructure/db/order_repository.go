package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/jnates/crud_golang/internal/domain/ports"
	"github.com/jnates/crud_golang/internal/infrastructure/db/queries"
)

type OrdersRepository struct {
	db *sql.DB
}

func NewOrdersRepository(db *sql.DB) ports.OrdersRepository {
	return &OrdersRepository{db: db}
}

func (r *OrdersRepository) ListOrders(ctx context.Context, status string) ([]model.Order, error) {
	var rows *sql.Rows
	var err error

	if status != "" {
		println("DEBUG: Executing QueryListOrdersFiltered with status:", status)
		rows, err = r.db.QueryContext(ctx, queries.QueryListOrdersFiltered, status)
	} else {
		println("DEBUG: Executing QueryListOrders (unfiltered)")
		rows, err = r.db.QueryContext(ctx, queries.QueryListOrders)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var o model.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.CustomerName, &o.CustomerEmail, &o.TotalAmount, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, err
		}
		o.OrderNumber = "ORD-" + o.ID.String()[:8]
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *OrdersRepository) ListOrdersByUserID(ctx context.Context, userID string, status string) ([]model.Order, error) {
	var rows *sql.Rows
	var err error

	if status != "" {
		println("DEBUG: Executing QueryListOrdersByUserIDFiltered - User:", userID, "Status:", status)
		rows, err = r.db.QueryContext(ctx, queries.QueryListOrdersByUserIDFiltered, userID, status)
	} else {
		println("DEBUG: Executing QueryListOrdersByUserID - User:", userID)
		rows, err = r.db.QueryContext(ctx, queries.QueryListOrdersByUserID, userID)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var o model.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.CustomerName, &o.CustomerEmail, &o.TotalAmount, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, err
		}
		o.OrderNumber = "ORD-" + o.ID.String()[:8]
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *OrdersRepository) ListOrderItems(ctx context.Context) ([]model.OrderItem, error) {
	rows, err := r.db.QueryContext(ctx, queries.QueryListOrderItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.OrderItem
	for rows.Next() {
		var i model.OrderItem
		var createdAt time.Time // Placeholder for scanning
		if err := rows.Scan(&i.ID, &i.OrderID, &i.ProductID, &i.Quantity, &i.PriceAtTime, &createdAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func (r *OrdersRepository) GetOrderByID(ctx context.Context, id string) (*model.Order, error) {
	var o model.Order
	err := r.db.QueryRowContext(ctx, queries.QueryGetOrderByID, id).Scan(
		&o.ID, &o.UserID, &o.CustomerName, &o.CustomerEmail, &o.TotalAmount, &o.Status, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		return nil, err
	}
	o.OrderNumber = "ORD-" + o.ID.String()[:8]

	// Fetch items
	rows, err := r.db.QueryContext(ctx, queries.QueryGetOrderItemsByOrderID, id)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var i model.OrderItem
			if err := rows.Scan(&i.ID, &i.OrderID, &i.ProductID, &i.ProductName, &i.Quantity, &i.PriceAtTime); err == nil {
				o.Items = append(o.Items, i)
			}
		}
	}

	return &o, nil
}

func (r *OrdersRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, queries.QueryCreateOrder,
		order.ID, order.UserID, order.TotalAmount, order.Status, order.CreatedAt, order.UpdatedAt)
	if err != nil {
		return err
	}

	for i := range order.Items {
		item := &order.Items[i]
		if item.ID == uuid.Nil {
			item.ID = uuid.New()
		}
		_, err = tx.ExecContext(ctx, queries.QueryCreateOrderItem,
			item.ID, order.ID, item.ProductID, item.Quantity, item.PriceAtTime, order.CreatedAt)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *OrdersRepository) UpdateOrder(ctx context.Context, order *model.Order) error {
	_, err := r.db.ExecContext(ctx, queries.QueryUpdateOrder,
		order.Status, order.TotalAmount, order.UpdatedAt, order.ID)
	return err
}

func (r *OrdersRepository) DeleteOrder(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, queries.QueryDeleteOrder, id)
	return err
}
