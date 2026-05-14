package infrastructure

import (
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger(debug bool) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func parseBool(boolString string) bool {
	if boolString == "" {
		return false // Valor por defecto silencioso
	}
	boolVal, err := strconv.ParseBool(boolString)
	if err != nil {
		log.Warn().Msgf("Invalid boolean value '%s', defaulting to false", boolString)
		return false
	}
	return boolVal
}
