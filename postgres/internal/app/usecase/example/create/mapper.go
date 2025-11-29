package create

import "github.com/wisaitas/db-sharding/postgres/internal/app/domain/entity"

func mapRequestToEntity(req Request) []entity.Example {
	return []entity.Example{
		{
			Name: req.Name,
		},
	}
}
