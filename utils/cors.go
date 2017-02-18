package utils

import "github.com/rs/cors"

func EnableCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"http://12*"},
	})
}
