package entities

import "gorm.io/gorm"

type Setting struct {
	ID    string
	Key   string
	Value string
}
type Settings struct {
	gorm.Model
	ActiveThemeID string `gorm:"column:active_theme_id"`
}
