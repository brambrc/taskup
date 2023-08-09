package Model

import (
	"fmt"
	"taskup/Database"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskName  string `json:"task_name"`
	Deadline  string `json:"deadline"`
	Status    string `json:"status"`
	IdProject string `json:"id_project"`
}

func (t *Task) SaveTask() (*Task, error) {
	var err error
	err = Database.Database.Create(&t).Error
	if err != nil {
		return &Task{}, err
	}
	return t, nil
}

func GetTasksByIdProject(id string) (*[]Task, error) {
	var tasks []Task
	err := Database.Database.Where("id_project = ?", id).Find(&tasks).Error
	if err != nil {
		return &[]Task{}, err
	}
	return &tasks, nil
}

func GetTasksById(id uint) (*Task, error) {
	var task Task
	err := Database.Database.Where("id = ?", id).Find(&task).Error
	if err != nil {
		return &Task{}, err
	}
	return &task, nil
}

func (u *Task) UpdateTask(id uint) (*Task, error) {

	var err error
	err = Database.Database.Model(&Task{}).Where("id = ?", id).Updates(&Task{
		TaskName: u.TaskName,
		Deadline: u.Deadline,
		Status:   u.Status,
	}).Error
	if err != nil {
		return &Task{}, err
	}
	return u, nil
}

func DeleteTask(id uint) (*Task, error) {
	var task Task
	if err := Database.Database.First(&task, id).Error; err != nil {
		return nil, err
	}

	if err := Database.Database.Delete(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func DeleteTaskMultiple(id string) error {

	if err := Database.Database.Where("id_project = ?", id).Delete(&Task{}).Error; err != nil {
		return err
	}

	return nil

}

func CountTotalTaskById(id string) int64 {
	var total int64
	fmt.Println("id project: ", id)
	Database.Database.Model(&Task{}).Where("id_project = ?", id).Count(&total)

	fmt.Println("total task: ", total)
	return total
}

func CountFinishedTask(id string) int64 {
	var total int64
	Database.Database.Model(&Task{}).Where("id_project = ? AND status = ?", id, "Selesai").Count(&total)
	return total
}
