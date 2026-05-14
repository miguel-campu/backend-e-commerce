package payment

import (
	"fmt"

	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/jnates/crud_golang/internal/domain/ports"
)

type MockPaymentAdapter struct{}

func NewMockAdapter() ports.PaymentProvider {
	return &MockPaymentAdapter{}
}

func (m *MockPaymentAdapter) GeneratePaymentLink(order *model.Order, email string) (*ports.PaymentResponse, error) {
	// Simulamos una respuesta de una pasarela
	fakeURL := fmt.Sprintf("https://mercadopago.com.co/test-pay?id=%s&total=%.2f", order.ID.String(), order.TotalAmount)

	return &ports.PaymentResponse{
		CheckoutURL:   fakeURL,
		TransactionID: "TRANS-MOCK-999",
	}, nil
}
