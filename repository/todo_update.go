package repository

import (
	"go-simple-template/model"
)

func (r *repository) ToDoUpdate(id string, todo model.Todo) error {
	return r.db.Model(&model.Todo{}).Where("id = ?", id).Updates(&todo).Error
}
