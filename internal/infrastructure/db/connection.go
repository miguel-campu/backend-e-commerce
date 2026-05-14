package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jnates/crud_golang/internal/infrastructure/kit/enum"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewPostgresConnection() *sql.DB {

	os.Setenv(enum.DBHost, "localhost")
	os.Setenv(enum.DBPort, "5432")
	os.Setenv(enum.DBUser, "postgres")
	os.Setenv(enum.DBPassword, "123456789")
	os.Setenv(enum.DBName, "tiendita")
	os.Setenv(enum.SSLMode, "disable")
	host := os.Getenv(enum.DBHost)
	port := os.Getenv(enum.DBPort)
	user := os.Getenv(enum.DBUser)
	password := os.Getenv(enum.DBPassword)
	dbname := os.Getenv(enum.DBName)
	sslmode := os.Getenv(enum.SSLMode)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	log.Debug().Str("dsn", dsn).Msg("Building Database connection")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening Database connection")
	}

	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Error connecting to the database")
	}

	log.Info().Msg("Connection to PostgreSQL established")
	return db
}
