package service

import (
	"go-simple-template/dto"
	"go-simple-template/model"
	"go-simple-template/pkg/logger"
	"go-simple-template/repository"
)

type service struct {
	repo repository.RepositoryInterface
}

var (
	logService = logger.NewLogger().Logger.With().Str("pkg", "service").Logger()
)

func NewService(repo repository.RepositoryInterface) *service {
	return &service{
		repo: repo,
	}
}

type ServiceInterface interface {
	Ping() error
	ToDoCreate(dataReq dto.ToDoCreateReq) error
	ToDoDelete(id int64) error
	ToDoAll() ([]model.Todo, error)
	ToDoUpdate(id string, dataReq dto.ToDoUpdateReq) error
}
