package create

import (
	"github.com/gofiber/fiber/v3"
	"github.com/wisaitas/db-sharding/postgres/internal/app/domain/repositoryshard1"
)

type Service interface {
	CreateExample(ctx fiber.Ctx, req Request) error
}

type service struct {
	exampleRepo repositoryshard1.ExampleRepository
}

func newService(
	exampleRepo repositoryshard1.ExampleRepository,
) Service {
	return &service{
		exampleRepo: exampleRepo,
	}
}

func (s *service) CreateExample(c fiber.Ctx, req Request) error {
	entities := mapRequestToEntity(req)

	if err := s.exampleRepo.Create(c, entities); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}
