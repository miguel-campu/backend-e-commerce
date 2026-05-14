package queries

const (
	QueryGetOrderByID = `
		SELECT o.id, o.user_id, u.full_name, u.email, o.total_amount, o.status, o.created_at, o.updated_at
		FROM orders o
		JOIN users u ON o.user_id = u.id
		WHERE o.id = $1
	`
	QueryCreateOrder = `
		INSERT INTO orders (id, user_id, total_amount, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	QueryUpdateOrder = `
		UPDATE orders
		SET status = $1, total_amount = $2, updated_at = $3
		WHERE id = $4
	`
	QueryDeleteOrder = `
		DELETE FROM orders
		WHERE id = $1
	`
	QueryListOrders = `
		SELECT o.id, o.user_id, u.full_name, u.email, o.total_amount, o.status, o.created_at, o.updated_at
		FROM orders o
		JOIN users u ON o.user_id = u.id
		ORDER BY o.created_at DESC
	`
	QueryListOrdersFiltered = `
		SELECT o.id, o.user_id, u.full_name, u.email, o.total_amount, o.status, o.created_at, o.updated_at
		FROM orders o
		JOIN users u ON o.user_id = u.id
		WHERE o.status = $1
		ORDER BY o.created_at DESC
	`
	QueryListOrdersByUserID = `
		SELECT o.id, o.user_id, u.full_name, u.email, o.total_amount, o.status, o.created_at, o.updated_at
		FROM orders o
		JOIN users u ON o.user_id = u.id
		WHERE o.user_id = $1
		ORDER BY o.created_at DESC
	`
	QueryListOrdersByUserIDFiltered = `
		SELECT o.id, o.user_id, u.full_name, u.email, o.total_amount, o.status, o.created_at, o.updated_at
		FROM orders o
		JOIN users u ON o.user_id = u.id
		WHERE o.user_id = $1 AND o.status = $2
		ORDER BY o.created_at DESC
	`
	QueryListOrderItems = `
		SELECT id, order_id, product_id, quantity, price_at_time, created_at
		FROM order_items
	`
	QueryGetOrderItemsByOrderID = `
		SELECT oi.id, oi.order_id, oi.product_id, p.name, oi.quantity, oi.price_at_time
		FROM order_items oi
		JOIN products p ON oi.product_id = p.id
		WHERE oi.order_id = $1
	`
	QueryCreateOrderItem = `
		INSERT INTO order_items (id, order_id, product_id, quantity, price_at_time, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
)
