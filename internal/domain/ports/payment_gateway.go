package ports

import "github.com/jnates/crud_golang/internal/domain/model"

// PaymentResponse contiene lo que necesitamos para redirigir al usuario
type PaymentResponse struct {
	CheckoutURL   string // URL a la que mandamos al cliente (PSE o Tarjeta)
	TransactionID string // ID que nos da la pasarela
}

// PaymentProvider es la interfaz agnóstica
type PaymentProvider interface {
	GeneratePaymentLink(order *model.Order, userEmail string) (*PaymentResponse, error)
}
