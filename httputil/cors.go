package httputil

import (
	"github.com/toshim45/jajak/config"

	"github.com/rs/cors"
)

func EnableCors(c config.Environment) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"http://12*"},
	})
}
