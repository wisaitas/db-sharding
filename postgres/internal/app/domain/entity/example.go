package entity

import "github.com/wisaitas/db-sharding/postgres/common"

type Example struct {
	common.Entity
	Name string `gorm:"column:name;type:varchar(255);not null"`
}
