package dbutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildDynamicQuery(t *testing.T) {
	baseQuery := "SELECT * FROM products"
	filters := map[string]interface{}{
		"name": "phone",
	}

	t.Run("With filters", func(t *testing.T) {
		query, args := BuildDynamicQuery(baseQuery, filters, 1)
		assert.Contains(t, query, "WHERE name ILIKE $1")
		assert.Equal(t, []interface{}{"%phone%"}, args)
	})

	t.Run("Without filters", func(t *testing.T) {
		query, args := BuildDynamicQuery(baseQuery, nil, 1)
		assert.Equal(t, baseQuery, query)
		assert.Empty(t, args)
	})
}

func TestAddPagination(t *testing.T) {
	query := "SELECT * FROM products"
	args := []interface{}{}
	
	t.Run("Add pagination", func(t *testing.T) {
		newQuery, newArgs := AddPagination(query, args, 1, 10, 0)
		assert.Contains(t, newQuery, "LIMIT $1 OFFSET $2")
		assert.Equal(t, []interface{}{10, 0}, newArgs)
	})
}
