package main

import (
	"flag"
	"os"

	"github.com/jnates/crud_golang/internal/infrastructure"
	http "github.com/jnates/crud_golang/internal/infrastructure/http"
	"github.com/jnates/crud_golang/internal/infrastructure/kit/enum"

	"github.com/jnates/crud_golang/internal/infrastructure/db"

	//"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	//"github.com/rs/zerolog/log"
)

// @title CRUD Golang API
// @version 1.0
// @description This is a sample server for managing users.
// @host localhost:8081
// @BasePath /

func main() {
	// 1. Flags (SIEMPRE primero)
	debug := flag.Bool("debug", false, "enable debug mode")
	flag.Parse()

	// 2. Variables de entorno
	//if err := godotenv.Load(); err != nil {
	//	log.Warn().Msg("No .env file found, using system environment variables")
	//}

	// 3. Logger
	infrastructure.InitLogger(*debug)

	// 4. Base de datos
	dbConn := db.NewPostgresConnection()
	defer dbConn.Close()

	// 5. Puerto
	port := os.Getenv(enum.APIPort)
	if port == "" {
		port = "8081"
	}

	// 6. Servidor
	http.Start(port, dbConn)
}
