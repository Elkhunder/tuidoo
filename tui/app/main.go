package app

import (
	"fmt"
	"log"
	"os"
	"tuidoo/managers"
	"tuidoo/services"
	"tuidoo/tui"

	tea "github.com/charmbracelet/bubbletea"
)

// RunTUI starts the TUI application
func RunTUI() error {
	// Initialize services
	sc, err := services.NewServiceCollection()
	if err != nil {
		return fmt.Errorf("failed to initialize services: %w", err)
	}
	defer sc.Close()

	// Initialize theme manager
	themeManager := managers.NewThemeManager()

	// Load user's theme preference
	settings, err := sc.SettingsService.GetAllSettings()
	if err == nil && settings.ActiveThemeID != "" {
		if err := themeManager.SetTheme(settings.ActiveThemeID); err != nil {
			log.Printf("Failed to load theme '%s', using default: %v", settings.ActiveThemeID, err)
		}
	}

	// Create Bubble Tea program
	m := tui.NewModel(sc, themeManager)

	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running tuidoo: %w", err)
	}

	return nil
}

// SeedDatabase seeds the database with sample data
func SeedDatabase() {
	sc, err := services.NewServiceCollection()
	if err != nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}
	defer sc.Close()

	if err := services.Seed(sc.DbService); err != nil {
		log.Fatalf("❌ Seeding failed: %v", err)
	}

	fmt.Println("✅ Database seeded successfully")
	os.Exit(0)
}

// ResetDatabase cleans and reseeds the database
func ResetDatabase() {
	sc, err := services.NewServiceCollection()
	if err != nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}
	defer sc.Close()

	fmt.Println("⚠️  This will delete all existing data. Continue? (y/N)")
	var response string
	fmt.Scanln(&response)

	if response != "y" && response != "Y" {
		fmt.Println("Reset cancelled")
		os.Exit(0)
	}

	if err := services.ResetAndSeed(sc.DbService); err != nil {
		log.Fatalf("❌ Reset failed: %v", err)
	}

	fmt.Println("✅ Database reset successfully")
	os.Exit(0)
}

// CleanDatabase removes all data from the database
func CleanDatabase() {
	sc, err := services.NewServiceCollection()
	if err != nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}
	defer sc.Close()

	fmt.Println("⚠️  This will delete all data. Continue? (y/N)")
	var response string
	fmt.Scanln(&response)

	if response != "y" && response != "Y" {
		fmt.Println("Clean cancelled")
		os.Exit(0)
	}

	if err := services.CleanDatabase(sc.DbService); err != nil {
		log.Fatalf("❌ Clean failed: %v", err)
	}

	fmt.Println("✅ Database cleaned successfully")
	os.Exit(0)
}
