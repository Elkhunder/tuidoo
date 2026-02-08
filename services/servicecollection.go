package services

import (
	"fmt"
	"log"
)

type ServiceCollection struct {
	DbService       *DbService
	SettingsService *SettingsService
	ThemeService    *ThemeService
	ToDoService     *ToDoService
	ProjectService  *ProjectService
	ToDoListService *ToDoListService
}

func NewServiceCollection() (*ServiceCollection, error) {
	sc := &ServiceCollection{}

	if err := sc.Init(); err != nil {
		return nil, err
	}

	return sc, nil
}

func (sc *ServiceCollection) Init() error {
	log.Println("Initializing services...")

	// 1. Database
	sc.DbService = Db
	if err := sc.DbService.Connect(); err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}

	// 2. Settings
	settingsService, err := NewSettingsService(sc.DbService)
	if err != nil {
		return fmt.Errorf("settings service initialization failed: %w", err)
	}
	sc.SettingsService = settingsService

	// 3. Theme
	sc.ThemeService = NewThemeService(sc.DbService, sc.SettingsService)

	// 4. Domain services
	sc.ToDoService = NewToDoService(sc.DbService)
	sc.ProjectService = NewProjectService(sc.DbService)
	sc.ToDoListService = NewToDoListService(sc.DbService)

	// 5. Seed
	if err := Seed(sc.DbService); err != nil {
		log.Printf("⚠️  Seeding failed (non-fatal): %v", err)
	}

	log.Println("✅ Services initialized successfully")
	return nil
}

func (sc *ServiceCollection) Close() error {
	log.Println("Shutting down services...")

	if sc.DbService != nil {
		if err := sc.DbService.Close(); err != nil {
			return fmt.Errorf("database close failed: %w", err)
		}
	}

	log.Println("✅ Services shut down successfully")
	return nil
}

// Reset cleans and reseeds the database
func (sc *ServiceCollection) Reset() error {
	log.Println("Resetting database...")

	if err := ResetAndSeed(sc.DbService); err != nil {
		return fmt.Errorf("reset failed: %w", err)
	}

	// Reload services after reset
	if err := sc.SettingsService.Reload(); err != nil {
		return fmt.Errorf("failed to reload settings after reset: %w", err)
	}

	log.Println("✅ Database reset complete")
	return nil
}

// HealthCheck verifies all services are operational
func (sc *ServiceCollection) HealthCheck() error {
	// Check database
	if sc.DbService == nil || sc.DbService.GetDB() == nil {
		return fmt.Errorf("database service not initialized")
	}

	ctx, cancel := sc.DbService.NewContext()
	defer cancel()

	if err := sc.DbService.GetDB().WithContext(ctx).Exec("SELECT 1").Error; err != nil {
		return fmt.Errorf("database not responding: %w", err)
	}

	// Check settings service
	if sc.SettingsService == nil {
		return fmt.Errorf("settings service not initialized")
	}

	if _, err := sc.SettingsService.GetAllSettings(); err != nil {
		return fmt.Errorf("settings service not operational: %w", err)
	}

	// Check theme service
	if sc.ThemeService == nil {
		return fmt.Errorf("theme service not initialized")
	}

	// Check todo service
	if sc.ToDoService == nil {
		return fmt.Errorf("todo service not initialized")
	}

	// Check project service
	if sc.ProjectService == nil {
		return fmt.Errorf("project service not initialized")
	}

	// Check todo list service
	if sc.ToDoListService == nil {
		return fmt.Errorf("todo list service not initialized")
	}

	return nil
}
