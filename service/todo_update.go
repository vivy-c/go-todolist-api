package service

import (
	"errors"
	"go-simple-template/dto"
	"go-simple-template/model"
)

func (s *service) ToDoUpdate(id string, dataReq dto.ToDoUpdateReq) error {
	todo := model.Todo{
		Todo:   dataReq.Todo,
		Title:  dataReq.Title,
		Detail: dataReq.Detail,
	}

	err := s.repo.ToDoUpdate(id, todo)
	if err != nil {
		return errors.New("failed to update book")
	}

	return nil
}
