package entities

import (
	"time"
	"tuidoo/enums"

	"gorm.io/gorm"
)

type ToDo struct {
	gorm.Model
	ProjectID   uint
	Project     Project
	ToDoListID  uint
	ToDoList    ToDoList
	Name        string
	Description *string
	Details     *string
	Priority    enums.Priority
	Status      enums.Status
	Color       string
	Done        bool
	DueDate     *time.Time
}
