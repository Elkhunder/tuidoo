package services

import (
	"fmt"

	e "tuidoo/entities"

	"gorm.io/gorm"
)

type SettingsService struct {
	db       *DbService
	settings *e.Settings
}

func NewSettingsService(dbService *DbService) (*SettingsService, error) {
	ss := &SettingsService{db: dbService}
	if err := ss.loadSettings(); err != nil {
		return nil, fmt.Errorf("failed to load settings: %w", err)
	}
	return ss, nil
}

func (ss *SettingsService) loadSettings() error {
	ctx, cancel := ss.db.NewContext()
	defer cancel()

	var settings e.Settings
	err := ss.db.GetDB().WithContext(ctx).First(&settings).Error

	if err != nil {
		// If no settings found, create default settings
		if err == gorm.ErrRecordNotFound {
			defaultSettings := e.Settings{
				ActiveThemeID: "dark",
			}

			if createErr := ss.db.GetDB().WithContext(ctx).Create(&defaultSettings).Error; createErr != nil {
				return fmt.Errorf("failed to create default settings: %w", createErr)
			}

			ss.settings = &defaultSettings
			return nil
		}
		return err
	}

	ss.settings = &settings
	return nil
}

func (ss *SettingsService) GetAllSettings() (*e.Settings, error) {
	if ss.settings == nil {
		if err := ss.loadSettings(); err != nil {
			return nil, err
		}
	}
	return ss.settings, nil
}

func (ss *SettingsService) Reload() error {
	return ss.loadSettings()
}

func (ss *SettingsService) GetActiveTheme() (string, error) {
	settings, err := ss.GetAllSettings()
	if err != nil {
		return "", fmt.Errorf("failed to get settings: %w", err)
	}
	return settings.ActiveThemeID, nil
}

func (ss *SettingsService) SetActiveTheme(themeId string) error {
	settings, err := ss.GetAllSettings()
	if err != nil {
		return fmt.Errorf("failed to get settings: %w", err)
	}

	ctx, cancel := ss.db.NewContext()
	defer cancel()

	settings.ActiveThemeID = themeId
	if err := ss.db.GetDB().WithContext(ctx).Save(&settings).Error; err != nil {
		return fmt.Errorf("failed to save theme: %w", err)
	}

	ss.settings = settings
	return nil
}
