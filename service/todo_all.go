package service

import (
	"go-simple-template/model"
)

func (s *service) ToDoAll() ([]model.Todo, error) {
	ToDo, err := s.repo.ToDoAll()
	if err != nil {
		return nil, err
	}

	return ToDo, nil
}
