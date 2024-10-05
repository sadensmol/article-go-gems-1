package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/codingconcepts/env"
)

type Config struct {
	Postgres
}

type Postgres struct {
	Host               string `env:"POSTGRES_HOST"`
	Port               int    `env:"POSTGRES_PORT"`
	User               string `env:"POSTGRES_USER"`
	Password           string `env:"POSTGRES_PASSWORD"`
	DBName             string `env:"POSTGRES_DB_NAME"`
	SchemaName         string `env:"POSTGRES_DB_SCHEMA_NAME"`
	DBSSLMode          string `env:"POSTGRES_DB_SSL_MODE"`
	MaxConnections     int    `env:"POSTGRES_MAX_CONNECTIONS"`
	MaxIdleConnections int    `env:"POSTGRES_MAX_IDLE_CONNECTIONS"`
}

func (c Postgres) ConnectionURL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s&search_path=%s", c.User, c.Password, c.Host, c.Port, c.DBName, c.DBSSLMode, c.SchemaName)
}

var once sync.Once
var instance Config

func GetConfig() *Config {
	once.Do(func() {
		if err := env.Set(&instance.Postgres); err != nil {
			log.Fatalf("cannot init config %s", err)
		}
	})

	return &instance
}
