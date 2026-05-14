package queries

const (
	QueryCreateUser = `
		INSERT INTO users (id, email, password_hash, role, full_name, avatar_url, google_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	QueryGetUserByID = `
		SELECT id, email, password_hash, role, full_name, COALESCE(avatar_url, ''), COALESCE(google_id, ''), created_at, updated_at
		FROM users WHERE id = $1
	`
	QueryGetUserByEmail = `
		SELECT id, email, password_hash, role, full_name, COALESCE(avatar_url, ''), COALESCE(google_id, ''), created_at, updated_at
		FROM users WHERE email = $1
	`
	QueryUpdateUser = `
		UPDATE users SET email = $1, role = $2, full_name = $3, avatar_url = $4, google_id = $5, updated_at = $6
		WHERE id = $7
	`
	QueryDeleteUser = `
		DELETE FROM users WHERE id = $1
	`
	QueryUpdateUserPassword = `
		UPDATE users SET password_hash = $1, updated_at = $2
		WHERE id = $3
	`
	QueryListUsers = `
		SELECT id, email, password_hash, role, full_name, COALESCE(avatar_url, ''), COALESCE(google_id, ''), created_at, updated_at
		FROM users
	`
)
