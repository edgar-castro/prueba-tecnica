package models

import (
	"go-test-api/database"
)

type Project struct {
	ID          int
	UserId      int
	Name        string
	Description string
	StartDate   string
	EndDate     string
	Budget      float32
}

type ProjectInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
	Budget      float32 `json:"budget"`
}

func CreateProject(project *Project) error {
	// return database.DB.Create(&project).Error
	return database.DB.Exec(
		"INSERT INTO project (user_id, name, description, start_date, end_date, budget) VALUES (?, ?, ?, ?, ?, ?)",
		project.UserId, project.Name, project.Description, project.StartDate, project.EndDate, project.Budget,
	).Error
}

func GetProjects(projects *[]Project, id int) error {
	return database.DB.Find(&projects, "id_user = ?", id).Error
}

func GetProject(project *Project, id int) error {
	return database.DB.Find(&project, id).Error
}

func UpdateProject(project *ProjectInput, id int) error {
	var recover Project
	database.DB.Find(&recover, id)
	recover.Name = project.Name
	recover.Description = project.Description
	recover.StartDate = project.StartDate
	recover.EndDate = project.EndDate
	recover.Budget = project.Budget
	return database.DB.Save(&recover).Error
}

func DeleteProject(id int) error {
	return database.DB.Delete(&Project{}, id).Error
}
