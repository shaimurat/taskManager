package postgres

import (
	"gorm.io/gorm"
	"log"
	"taskManager/internal/repository/postgres/models"
)

type TaskRepoSQL struct {
	DB *gorm.DB
}

func NewTaskRepoSQL(db *gorm.DB) *TaskRepoSQL {
	return &TaskRepoSQL{
		DB: db,
	}
}

func (r *TaskRepoSQL) Create(task models.TaskDoc) error {
	result := r.DB.Create(&task)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Rows affected: ", result.RowsAffected)
	log.Println("Task created successfully:", task)
	return nil
}

func (r *TaskRepoSQL) Update(changes interface{}, task models.TaskDoc) error {
	result := r.DB.Model(&task).Where("user_id = ?", task.ID).Updates(changes)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Rows affected: ", result.RowsAffected)
	log.Println("Task updated successfully:", changes)
	return nil
}
func (r *TaskRepoSQL) Delete(task models.TaskDoc) error {
	result := r.DB.Delete(&task)
	if result.Error != nil {
		return result.Error
	}
	log.Println("Rows affected: ", result.RowsAffected)
	log.Println("Task deleted successfully:", task)
	return nil
}
