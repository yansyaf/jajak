package utils

import "github.com/rs/cors"

func EnableCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"http://128.199.91.72", "http://127.0.0.1"},
	})
}
