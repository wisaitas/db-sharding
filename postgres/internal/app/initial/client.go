package initial

import (
	"log"

	"github.com/wisaitas/db-sharding/pkg/db/postgres"
	"github.com/wisaitas/db-sharding/postgres/internal/app"
	"gorm.io/gorm"
)

type client struct {
	postgresShard1 *gorm.DB
	postgresShard2 *gorm.DB
}

func newClient() *client {
	db1, err := postgres.NewPostgreSQL(app.Config.Shard0.Postgres)
	if err != nil {
		log.Fatalf("error creating postgres shard 0 client: %v", err)
	}

	db2, err := postgres.NewPostgreSQL(app.Config.Shard1.Postgres)
	if err != nil {
		log.Fatalf("error creating postgres shard 1 client: %v", err)
	}

	return &client{
		postgresShard1: db1,
		postgresShard2: db2,
	}
}
