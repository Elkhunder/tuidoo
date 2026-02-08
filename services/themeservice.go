package services

import (
	"context"
	"fmt"
	"time"

	e "tuidoo/entities"
)

type ThemeService struct {
	db              *DbService
	settingsService *SettingsService
}

func NewThemeService(dbService *DbService, settingsService *SettingsService) *ThemeService {
	return &ThemeService{
		db:              dbService,
		settingsService: settingsService,
	}
}

// GetCurrentTheme retrieves the active theme based on settings
func (ts *ThemeService) GetCurrentTheme() (*e.Theme, error) {
	themeID, err := ts.settingsService.GetActiveTheme()
	if err != nil {
		return nil, fmt.Errorf("failed to get active theme ID: %w", err)
	}

	ctx, cancel := ts.db.NewContext()
	defer cancel()

	var theme e.Theme
	if err := ts.db.GetDB().WithContext(ctx).Where("id = ?", themeID).First(&theme).Error; err != nil {
		// Return default theme if not found
		return ts.getDefaultTheme(), nil
	}

	return &theme, nil
}

// SaveCurrentTheme saves or updates a theme
func (ts *ThemeService) SaveCurrentTheme(theme *e.Theme) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ts.db.GetDB().WithContext(ctx).Save(theme).Error; err != nil {
		return fmt.Errorf("failed to save theme: %w", err)
	}

	return nil
}

// GetThemeByID retrieves a specific theme by ID
func (ts *ThemeService) GetThemeByID(themeID string) (*e.Theme, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	var theme e.Theme
	if err := ts.db.GetDB().WithContext(ctx).Where("id = ?", themeID).First(&theme).Error; err != nil {
		return nil, fmt.Errorf("theme not found: %w", err)
	}

	return &theme, nil
}

// ListThemes returns all available themes
func (ts *ThemeService) ListThemes() ([]e.Theme, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	var themes []e.Theme
	if err := ts.db.GetDB().WithContext(ctx).Find(&themes).Error; err != nil {
		return nil, fmt.Errorf("failed to list themes: %w", err)
	}

	return themes, nil
}

// getDefaultTheme returns a default theme when none is configured
func (ts *ThemeService) getDefaultTheme() *e.Theme {
	return &e.Theme{
		// Add default theme properties based on your entity structure
	}
}
