package common

import (
	"github.com/google/uuid"
)

type Entity struct {
	ID uuid.UUID `gorm:"type:uuid;primarykey"`
}
