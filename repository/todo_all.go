package repository

import (
	"go-simple-template/model"
)

func (r *repository) ToDoAll() ([]model.Todo, error) {
	var ToDo []model.Todo
	err := r.db.Find(&ToDo).Error
	if err != nil {
		return nil, err
	}

	return ToDo, nil
}
