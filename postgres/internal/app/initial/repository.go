package initial

import (
	appRepository "github.com/wisaitas/db-sharding/postgres/internal/app/domain/repositoryshard1"
)

type repository struct {
	ExampleRepository appRepository.ExampleRepository
}

func newRepository(client *client) *repository {
	return &repository{
		ExampleRepository: appRepository.NewExampleRepository(client.postgresShard1),
	}
}
