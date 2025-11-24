package postgres

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	First(ctx context.Context, conds ...interface{}) (*T, error)
	Find(ctx context.Context, conds ...interface{}) ([]T, error)
	Create(ctx context.Context, entity *T) error
	Updates(ctx context.Context, entity *T, values interface{}) error
	Delete(ctx context.Context, conds ...interface{}) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{
		db: db,
	}
}

func (r *baseRepository[T]) First(ctx context.Context, conds ...interface{}) (*T, error) {
	var dest T
	err := r.db.WithContext(ctx).First(&dest, conds...).Error
	if err != nil {
		return nil, err
	}
	return &dest, nil
}

func (r *baseRepository[T]) Find(ctx context.Context, conds ...interface{}) ([]T, error) {
	var dest []T
	err := r.db.WithContext(ctx).Find(&dest, conds...).Error
	if err != nil {
		return nil, err
	}
	return dest, nil
}

func (r *baseRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *baseRepository[T]) Updates(ctx context.Context, entity *T, values interface{}) error {
	return r.db.WithContext(ctx).Model(entity).Updates(values).Error
}

func (r *baseRepository[T]) Delete(ctx context.Context, conds ...interface{}) error {
	var entity T
	return r.db.WithContext(ctx).Delete(&entity, conds...).Error
}
