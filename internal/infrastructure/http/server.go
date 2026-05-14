package infrastructure

import (

	//falta la libreria de Nates
	"database/sql"

	"net/http"

	"github.com/jnates/crud_golang/internal/infrastructure/di"
	"github.com/jnates/crud_golang/internal/infrastructure/http/handler"
	validatorPackage "github.com/jnates/crud_golang/internal/infrastructure/http/validetor"
	"github.com/jnates/crud_golang/internal/infrastructure/kit/enum"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start(port string, db *sql.DB) {
	conn := db
	container := di.BuildContainer(conn)
	if container == nil {
		log.Fatal().Msg(" Error building DI container")
		return
	}

	err := container.Invoke(func(orderHandler *handler.OrderHandler, userHandler *handler.UserHandler, productHandler *handler.ProductHandler, analyticsHandler *handler.AnalyticsHandler) {
		e := echo.New()

		e.HideBanner = true
		e.Logger.SetOutput(log.Logger)

		e.Validator = validatorPackage.NewValidator()

		//conection with REACT JS

		// Swagger docs
		e.GET("/swagger/*", echoSwagger.WrapHandler)

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
			// Permitimos los métodos y los headers que usa Axios
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		}))

		// Rutas de Usuarios
		auth := e.Group("/auth")
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)

		users := e.Group("/users")
		users.GET("", userHandler.List)
		users.POST("", userHandler.Create)
		users.GET("/:id", userHandler.Get)
		users.PUT("/:id", userHandler.Update)
		users.PATCH("/:id/password", userHandler.UpdatePassword)
		users.DELETE("/:id", userHandler.Delete)

		// Rutas de Productos
		products := e.Group("/products")
		products.GET("", productHandler.List)
		products.GET("/:id", productHandler.Get)
		products.POST("", productHandler.Create)
		products.PUT("/:id", productHandler.Update)
		products.DELETE("/:id", productHandler.Delete)

		// Rutas de API
		orders := e.Group("/orders")
		orders.GET("", orderHandler.List)
		orders.GET("/:id", orderHandler.GetOrder)
		orders.POST("", orderHandler.Create)
		orders.PUT("/:id", orderHandler.Update)
		orders.DELETE("/:id", orderHandler.Delete)

		orders.POST("/:id/pay", orderHandler.HandlePayment)

		// Rutas de Analítica
		analytics := e.Group("/analytics")
		analytics.GET("/metrics", analyticsHandler.GetDashboardMetrics)
		analytics.GET("/revenue", analyticsHandler.GetRevenueChart)
		analytics.GET("/categories", analyticsHandler.GetCategorySales)
		analytics.GET("/top-products", analyticsHandler.GetTopProducts)

		log.Info().Str(enum.APIPort, port).Msg("Listened server")
		if err := e.Start(":" + port); err != nil {
			log.Fatal().Err(err).Msg("Error inicializing server")
		}
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Error inicializing dependencies with dig")
	}

}
