package Model

import (
	"strconv"
	"taskup/Database"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ProjectName  string `json:"project_name"`
	Descriptions string `json:"descriptions"`
	Category     string `json:"category"`
	Deadline     string `json:"deadline"`
	Priority     string `json:"priority"`
	IdUser       uint   `json:"id_user"`
}

type ResponseProject struct {
	ID           uint    `json:"id"`
	ProjectName  string  `json:"project_name"`
	Descriptions string  `json:"descriptions"`
	Category     string  `json:"category"`
	Deadline     string  `json:"deadline"`
	Priority     string  `json:"priority"`
	Percentage   float64 `json:"percentage"`
	TaskDone     int64   `json:"task_done"`
	TaskTotal    int64   `json:"task_total"`
	IdUser       uint    `json:"id_user"`
}

func (p *Project) SaveProject() (*Project, error) {
	var err error
	err = Database.Database.Create(&p).Error
	if err != nil {
		return &Project{}, err
	}
	return p, nil
}

func GetProjectsByUserID(id uint) ([]ResponseProject, error) {
	var projects []Project
	err := Database.Database.Where("id_user = ?", id).Find(&projects).Error

	if err != nil {
		return nil, err
	}

	var newProjects []ResponseProject
	for _, project := range projects {

		//convert id into string
		id := strconv.Itoa(int(project.ID))
		//count total tasks
		totalTasks := CountTotalTaskById(id)

		//count task that has been done

		doneTask := CountFinishedTask(id)

		var completionRatio float64
		if totalTasks > 0 {
			completionRatio = float64(doneTask) / float64(totalTasks) * 100
		}

		newProjects = append(newProjects, ResponseProject{
			ID:           project.ID,
			ProjectName:  project.ProjectName,
			Descriptions: project.Descriptions,
			Category:     project.Category,
			Deadline:     project.Deadline,
			Priority:     project.Priority,
			Percentage:   completionRatio,
			TaskDone:     int64(doneTask),
			TaskTotal:    int64(totalTasks),
			IdUser:       project.IdUser,
		})
	}

	return newProjects, nil
}

func GetProjectByID(id uint) (*Project, error) {
	var project Project
	err := Database.Database.Where("id = ?", id).Find(&project).Error
	if err != nil {
		return &Project{}, err
	}
	return &project, nil
}

func (p *Project) UpdateProject(id uint) (*Project, error) {
	var err error
	err = Database.Database.Where("id = ?", id).Updates(&p).Error
	if err != nil {
		return &Project{}, err
	}
	return p, nil
}

func DeleteProject(id uint) (*Project, error) {

	var project Project
	if err := Database.Database.First(&project, id).Error; err != nil {
		return &Project{}, err
	}
	if err := Database.Database.Delete(&project).Error; err != nil {
		return &Project{}, err
	}
	return &project, nil
}
