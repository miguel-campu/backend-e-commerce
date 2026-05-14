package dbutils

import (
	"database/sql"
	"fmt"
	"strings"
)

// BuildDynamicQuery construye un query base con filtros tipo ILIKE y placeholders tipo $1, $2...
func BuildDynamicQuery(baseQuery string, filters map[string]interface{}, startIndex int) (string, []interface{}) {
	var args []interface{}
	var conditions []string
	argPos := startIndex

	for key, val := range filters {
		conditions = append(conditions, fmt.Sprintf("%s ILIKE $%d", key, argPos))
		args = append(args, fmt.Sprintf("%%%v%%", val))
		argPos++
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	return baseQuery, args
}

// AddPagination agrega LIMIT y OFFSET con placeholders dinámicos
func AddPagination(query string, args []interface{}, startIndex int, limit, offset int) (string, []interface{}) {
	query += fmt.Sprintf(" ORDER BY id LIMIT $%d OFFSET $%d", startIndex, startIndex+1)
	args = append(args, limit, offset)
	return query, args
}

// ScanRows escanea múltiples filas y aplica una función personalizada
func ScanRows[T any](rows *sql.Rows, scanFn func(*sql.Rows) (*T, error)) ([]*T, error) {
	defer rows.Close()
	var results []*T
	for rows.Next() {
		item, err := scanFn(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, item)
	}
	return results, nil
}
