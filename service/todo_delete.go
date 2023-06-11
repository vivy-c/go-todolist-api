package service

import (
	"errors"
)

func (s *service) ToDoDelete(id int64) error {
	err := s.repo.ToDoDelete(id)
	if err != nil {
		return errors.New("failed to delete book")
	}

	return nil
}
