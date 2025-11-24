package app

import postgresql "github.com/wisaitas/db-sharding/pkg/db/postgres"

var Config struct {
	Server struct {
		Port string `env:"PORT" envDefault:"8080"`
	} `envPrefix:"SERVER_"`
	Shard0 struct {
		Postgres postgresql.Config
	} `envPrefix:"SHARD0_"`
	Shard1 struct {
		Postgres postgresql.Config
	} `envPrefix:"SHARD1_"`
}
