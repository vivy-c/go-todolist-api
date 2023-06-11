package service

import (
	"errors"
	"go-simple-template/dto"
	"go-simple-template/model"
)

func (s *service) ToDoCreate(dataReq dto.ToDoCreateReq) error {
	todo := model.Todo{
		Todo:   dataReq.Todo,
		Title:  dataReq.Title,
		Detail: dataReq.Detail,
	}

	err := s.repo.ToDoCreate(todo)
	if err != nil {
		return errors.New("failed create todo")
	}

	return nil
}
