package repository

import (
	"go-simple-template/model"
)

func (r *repository) ToDoDelete(id int64) error {
	return r.db.Delete(&model.Todo{}, id).Error
}
