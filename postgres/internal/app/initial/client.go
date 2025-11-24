package initial

import (
	"log"

	"github.com/wisaitas/db-sharding/pkg/db/postgres"
	"github.com/wisaitas/db-sharding/postgres/internal/app"
	"gorm.io/gorm"
)

type client struct {
	shards map[int]*gorm.DB
}

func newClient() *client {
	shards := make(map[int]*gorm.DB)

	db0, err := postgres.NewPostgreSQL(app.Config.Shard0.Postgres)
	if err != nil {
		log.Fatalf("error creating postgres shard 0 client: %v", err)
	}
	shards[0] = db0

	db1, err := postgres.NewPostgreSQL(app.Config.Shard1.Postgres)
	if err != nil {
		log.Fatalf("error creating postgres shard 1 client: %v", err)
	}
	shards[1] = db1

	return &client{
		shards: shards,
	}
}
