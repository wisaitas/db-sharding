package app

import postgresql "github.com/wisaitas/db-sharding/pkg/db/postgres"

var Config struct {
	Server struct {
		Port string `env:"PORT" envDefault:"8080"`
	} `envPrefix:"SERVER_"`
	Postgres postgresql.Config `envPrefix:"POSTGRES_"`
}
