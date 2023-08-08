package Controller

import (
	"net/http"
	"strconv"
	"taskup/Model"

	"github.com/gin-gonic/gin"
)

type Task struct {
	TaskName  string `json:"task_name"`
	Deadline  string `json:"deadline"`
	Status    string `json:"status"`
	IdProject string `json:"id_project"`
}

func CreateTask(context *gin.Context) {
	var Input Task
	if err := context.ShouldBindJSON(&Input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := Model.Task{
		TaskName:  Input.TaskName,
		Deadline:  Input.Deadline,
		Status:    Input.Status,
		IdProject: Input.IdProject,
	}

	input, err := task.SaveTask()

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "Task created successfully!", "task": input})

}

func GetTasksByIdProject(context *gin.Context) {

	id_project := context.Param("id_project")

	task, err := Model.GetTasksByIdProject(id_project)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "Task retrieved successfully!", "task": task})
}

func GetTasks(context *gin.Context) {

	id := context.Param("id")

	//convert id to uint
	convert_id, err := strconv.Atoi(id)
	unsignedNum := uint(convert_id)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	tasks, err := Model.GetTasksById(unsignedNum)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "Task retrieved successfully!", "tasks": tasks})
}

func UpdateTask(context *gin.Context) {

	id := context.Param("id")

	//convert id to uint
	convert_id, err := strconv.Atoi(id)
	unsignedNum := uint(convert_id)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	var Update Task

	if err := context.ShouldBindJSON(&Update); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := Model.Task{
		TaskName: Update.TaskName,
		Deadline: Update.Deadline,
		Status:   Update.Status,
	}

	_, err = task.UpdateTask(unsignedNum)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "Task updated successfully!"})

}

func DeleteTask(context *gin.Context) {

	id := context.Param("id")

	//convert id to uint
	convert_id, err := strconv.Atoi(id)
	unsignedNum := uint(convert_id)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	//delete task
	task, err := Model.DeleteTask(unsignedNum)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!", "task": task})

}
