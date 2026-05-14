package handler

import (
	"net/http"

	"github.com/jnates/crud_golang/internal/application"
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	Service    *application.OrderService
	PaymentSvc *application.PaymentService
}

func NewOrderHandler(svc *application.OrderService, paySvc *application.PaymentService) *OrderHandler {
	return &OrderHandler{
		Service:    svc,
		PaymentSvc: paySvc,
	}
}

func (h *OrderHandler) List(c echo.Context) error {
	userID := c.QueryParam("customerId")
	role := c.QueryParam("role")
	status := c.QueryParam("status")

	println("DEBUG: List Orders - UserID:", userID, "Role:", role, "Status:", status)

	isAdmin := role == "admin"

	if !isAdmin && userID == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"data":  []model.Order{},
			"total": 0,
		})
	}
	
	orders, err := h.Service.List(c.Request().Context(), userID, isAdmin, status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data":  orders,
		"total": len(orders),
	})
}

func (h *OrderHandler) GetOrder(c echo.Context) error {
	id := c.Param("id")
	order, err := h.Service.GetOrder(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "order not found"})
	}
	return c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Create(c echo.Context) error {
	var order model.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	if err := h.Service.Create(c.Request().Context(), &order); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var order model.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}
	order.ID.UnmarshalText([]byte(id))

	if err := h.Service.Update(c.Request().Context(), &order); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.Service.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *OrderHandler) HandlePayment(c echo.Context) error {
	id := c.Param("id")
	
	order, err := h.Service.GetOrder(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "order not found"})
	}

	order.Status = "paid"
	if err := h.Service.Update(c.Request().Context(), order); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "could not update order status"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "payment successful", "status": "paid"})
}
