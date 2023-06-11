package repository

import (
	"go-simple-template/model"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock *mock.Mock
}

func NewRepositoryMock(mock *mock.Mock) RepositoryInterface {
	return &RepositoryMock{Mock: mock}
}

func (r *RepositoryMock) Ping() error {
	args := r.Mock.Called()
	return args.Error(0)
}
func (r *RepositoryMock) ToDoCreate(todo model.Todo) error {
	args := r.Mock.Called()
	return args.Error(0)
}
func (r *RepositoryMock) ToDoDelete(id int64) error {
	args := r.Mock.Called()
	return args.Error(0)
}
func (r *RepositoryMock) ToDoAll() ([]model.Todo, error) {
	args := r.Mock.Called()
	return args.Get(0).([]model.Todo), args.Error(1)
}
func (r *RepositoryMock) ToDoUpdate(id string, todo model.Todo) error {
	args := r.Mock.Called()
	return args.Error(0)
}
