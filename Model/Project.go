package Model

import (
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

func (p *Project) SaveProject() (*Project, error) {
	var err error
	err = Database.Database.Create(&p).Error
	if err != nil {
		return &Project{}, err
	}
	return p, nil
}

func GetProjectsByUserID(id uint) ([]Project, error) {
	var projects []Project
	err := Database.Database.Where("id_user = ?", id).Find(&projects).Error
	if err != nil {
		return []Project{}, err
	}
	return projects, nil
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
