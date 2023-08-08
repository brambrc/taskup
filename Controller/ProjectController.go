package Controller

import (
	"net/http"
	"strconv"
	"taskup/Model"

	"github.com/gin-gonic/gin"
)

type Project struct {
	ProjectName  string `json:"project_name"`
	Descriptions string `json:"descriptions"`
	Category     string `json:"category"`
	Deadline     string `json:"deadline"`
	Priority     string `json:"priority"`
	IdUser       uint   `json:"id_user"`
}

func CreateProject(context *gin.Context) {
	var Input Project
	if err := context.ShouldBindJSON(&Input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token := context.Request.Header.Get("Authorization")

	id, err := Model.GetAuthenticatedID(token)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := Model.Project{
		ProjectName:  Input.ProjectName,
		Descriptions: Input.Descriptions,
		Category:     Input.Category,
		Deadline:     Input.Deadline,
		Priority:     Input.Priority,
		IdUser:       id,
	}

	input, err := project.SaveProject()

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "Project created successfully!", "project": input})

}

func GetProjects(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	id, err := Model.GetAuthenticatedID(token)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	projects, err := Model.GetProjectsByUserID(id)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Projects retrieved successfully!", "projects": projects})

}

func GetProjectByID(context *gin.Context) {

	id := context.Param("id")

	//convert id to uint
	convert_id, err := strconv.Atoi(id)
	unsignedNum := uint(convert_id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := Model.GetProjectByID(unsignedNum)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "Project retrieved successfully!", "project": project})
}

func UpdateProject(context *gin.Context) {

	id := context.Param("id")

	//convert id to uint
	convert_id, err := strconv.Atoi(id)
	unsignedNum := uint(convert_id)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Update Project
	if err := context.ShouldBindJSON(&Update); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	project := Model.Project{
		ProjectName:  Update.ProjectName,
		Descriptions: Update.Descriptions,
		Category:     Update.Category,
		Deadline:     Update.Deadline,
		Priority:     Update.Priority,
	}

	_, err = project.UpdateProject(unsignedNum)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Project updated successfully!"})

}

func DeleteProject(context *gin.Context) {

	id := context.Param("id")

	//convert id to uint
	convert_id, err := strconv.Atoi(id)
	unsignedNum := uint(convert_id)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	_, err = Model.DeleteTaskMultiple(id)

}
