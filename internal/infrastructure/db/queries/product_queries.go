package queries

const (
	QueryCreateProduct = `
		INSERT INTO products (id, name, description, price, stock, image_url, category, sku, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	QueryGetProductByID = `
		SELECT id, name, description, price, stock, COALESCE(image_url, ''), category, COALESCE(sku, ''), COALESCE(is_active, true), created_at, updated_at
		FROM products WHERE id = $1
	`
	QueryUpdateProduct = `
		UPDATE products SET name = $1, description = $2, price = $3, stock = $4, image_url = $5, category = $6, sku = $7, is_active = $8, updated_at = $9
		WHERE id = $10
	`
	QueryDeleteProduct = `
		DELETE FROM products WHERE id = $1
	`
	QueryListProducts = `
		SELECT id, name, description, price, stock, COALESCE(image_url, ''), category, COALESCE(sku, ''), COALESCE(is_active, true), created_at, updated_at
		FROM products
	`
)
