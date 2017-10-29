package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Port          string `envconfig:"JJ_PORT" default:"8071"`
	MongoHost     string `envconfig:"JJ_MONGO_HOST" default:"127.0.0.1"`
	MongoPort     string `envconfig:"JJ_MONGO_PORT" default:"27017"`
	MongoDBName   string `envconfig:"JJ_MONGO_DB_NAME" default:"jajak"`
	RedisHost     string `envconfig:"JJ_REDIS_HOST" default:"127.0.0.1"`
	RedisPort     string `envconfig:"JJ_REDIS_PORT" default:"6379"`
	EnableSwagger bool   `envconfig:"JJ_ENABLE_SWAGGER" default:"true"`
	AllowedOrigin string `envconfig:"JJ_ALLOWED_ORIGIN" default:"*"`
}

var env Environment
var once sync.Once

func NewEnv() Environment {
	once.Do(func() {
		err := envconfig.Process("", &env)
		if err != nil {
			log.Fatal("Can't load config: ", err)
		}
	})

	return env
}
