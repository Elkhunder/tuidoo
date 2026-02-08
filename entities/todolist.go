package entities

import "gorm.io/gorm"

type ToDoList struct {
	gorm.Model
	Name  string
	Color string
	ToDos []ToDo
}
