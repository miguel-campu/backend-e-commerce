package di

import (
	"database/sql"

	"github.com/jnates/crud_golang/internal/application"
	"github.com/jnates/crud_golang/internal/domain/ports"
	"github.com/jnates/crud_golang/internal/infrastructure/db"
	"github.com/jnates/crud_golang/internal/infrastructure/http/handler"
	"github.com/jnates/crud_golang/internal/infrastructure/payment" // Nuevo import
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

func BuildContainer(conn *sql.DB) *dig.Container {
	log.Debug().Msg("Starting dependency container construction")
	container := dig.New()

	// 1. Repositorios de Base de Datos
	container.Provide(func() ports.OrdersRepository {
		return db.NewOrdersRepository(conn)
	})
	container.Provide(func() ports.UserRepository {
		return db.NewUserRepository(conn)
	})
	container.Provide(func() ports.ProductRepository {
		return db.NewProductRepository(conn)
	})

	// 2. ADAPTERS DE INFRAESTRUCTURA (Aquí registras el Mock)
	// Cuando quieras usar PayU, solo cambias esta línea por payment.NewPayUAdapter
	container.Provide(func() ports.PaymentProvider {
		log.Debug().Msg("Recording Payment Gateway (Mock)")
		return payment.NewMockAdapter()
	})

	// 3. SERVICIOS DE APLICACIÓN
	container.Provide(func(repo ports.OrdersRepository) *application.OrderService {
		return application.NewOrderService(repo)
	})

	container.Provide(func(repo ports.UserRepository) *application.UserService {
		return application.NewUserService(repo)
	})
	container.Provide(func(repo ports.ProductRepository) *application.ProductService {
		return application.NewProductService(repo)
	})

	// Registramos el nuevo PaymentService
	container.Provide(func(repo ports.OrdersRepository, gw ports.PaymentProvider) *application.PaymentService {
		log.Debug().Msg("Recording PaymentService")
		return application.NewPaymentService(repo, gw)
	})

	// 4. HANDLERS (Controladores)
	// Actualizamos el OrderHandler para que ahora también reciba el PaymentService
	container.Provide(func(svc *application.OrderService, paySvc *application.PaymentService) *handler.OrderHandler {
		log.Debug().Msg("Recording OrderHandler with Payment capabilities")
		return handler.NewOrderHandler(svc, paySvc)
	})

	container.Provide(func(svc *application.UserService) *handler.UserHandler {
		return handler.NewUserHandler(svc)
	})
	container.Provide(func(svc *application.ProductService) *handler.ProductHandler {
		return handler.NewProductHandler(svc)
	})

	container.Provide(func(oSvc *application.OrderService, pSvc *application.ProductService, uSvc *application.UserService) *handler.AnalyticsHandler {
		return handler.NewAnalyticsHandler(oSvc, pSvc, uSvc)
	})

	log.Debug().Msg("Container built successfully")
	return container
}
