package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name  string
	Color string
	ToDos []ToDo
}
