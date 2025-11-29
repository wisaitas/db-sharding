package repositoryshard1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/wisaitas/db-sharding/postgres/internal/app/domain/entity"
	"gorm.io/gorm"
)

type ExampleRepository interface {
	Create(c fiber.Ctx, entities []entity.Example) error
}

type exampleRepository struct {
	postgresShard1 *gorm.DB
}

func NewExampleRepository(
	postgresShard1 *gorm.DB,
) ExampleRepository {
	return &exampleRepository{
		postgresShard1: postgresShard1,
	}
}

func (r *exampleRepository) Create(c fiber.Ctx, entities []entity.Example) error {
	return r.postgresShard1.WithContext(c.Context()).Create(&entities).Error
}
