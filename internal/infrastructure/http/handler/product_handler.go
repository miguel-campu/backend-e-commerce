package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jnates/crud_golang/internal/application"
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service *application.ProductService
}

func NewProductHandler(service *application.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Create(c echo.Context) error {
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request format"})
	}

	if err := c.Validate(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.service.CreateProduct(c.Request().Context(), &product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) Get(c echo.Context) error {
	id := c.Param("id")
	product, err := h.service.GetProduct(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request format"})
	}

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid product id"})
	}
	product.ID = parsedID

	if err := c.Validate(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.service.UpdateProduct(c.Request().Context(), &product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error: " + err.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteProduct(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *ProductHandler) List(c echo.Context) error {
	products, err := h.service.ListProducts(c.Request().Context())
	if (err != nil) {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  products,
		"total": len(products),
	})
}
