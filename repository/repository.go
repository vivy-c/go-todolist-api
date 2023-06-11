package repository

import (
	"go-simple-template/cache"
	"go-simple-template/model"
	"go-simple-template/pkg/logger"

	"gorm.io/gorm"
)

type repository struct {
	db    *gorm.DB
	cache *cache.Cache
}

var (
	logRepo = logger.NewLogger().Logger.With().Str("pkg", "repository").Logger()
)

func NewRepository(db *gorm.DB, cache *cache.Cache) *repository {
	return &repository{
		db:    db,
		cache: cache,
	}
}

type RepositoryInterface interface {
	Ping() error
	ToDoCreate(todo model.Todo) error
	ToDoDelete(id int64) error
	ToDoAll() ([]model.Todo, error)
	ToDoUpdate(id string, todo model.Todo) error
}
