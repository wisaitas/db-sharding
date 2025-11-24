package initial

import (
	"log"

	"github.com/wisaitas/db-sharding/pkg/db/postgres"
	"github.com/wisaitas/db-sharding/postgres/internal/app"
	"gorm.io/gorm"
)

type client struct {
	postgres *gorm.DB
}

func newClient() *client {
	dbClient, err := postgres.NewPostgreSQL(app.Config.Postgres)
	if err != nil {
		log.Fatalf("error creating postgres client: %v", err)
	}

	return &client{
		postgres: dbClient,
	}
}
