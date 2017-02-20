package utils

import (
	"github.com/rs/cors"
)

func EnableCors(c Config) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"http://12*"},
	})
}
