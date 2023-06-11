package repository

import (
	"go-simple-template/model"
)

func (r *repository) ToDoCreate(todo model.Todo) error {
	return r.db.Model(&model.Todo{}).Create(&todo).Error
}
