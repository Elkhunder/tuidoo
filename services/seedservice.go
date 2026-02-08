package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"tuidoo/entities"
	"tuidoo/enums"
	"tuidoo/generated"

	"gorm.io/gorm"
)

// Seed initializes all default data with transaction support
func Seed(dbService *DbService) error {
	log.Println("🌱 Starting database seeding...")

	db := dbService.GetDB()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Use transaction for atomic seeding
	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := seedSettings(tx, ctx); err != nil {
			return err
		}

		if err := seedProjects(tx, ctx); err != nil {
			return err
		}

		if err := seedToDoLists(tx, ctx); err != nil {
			return err
		}

		if err := seedToDos(tx, ctx); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("❌ Seeding failed: %v", err)
		return err
	}

	if err := verifySeeding(db, ctx); err != nil {
		log.Printf("⚠️ Warning: Verification failed: %v", err)
	}

	log.Println("🎉 Seed complete!")
	return nil
}

// CleanDatabase removes all existing data with error handling
func CleanDatabase(dbService *DbService) error {
	log.Println("🧹 Cleaning existing data...")

	db := dbService.GetDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Delete in reverse dependency order with error checking
	if err := db.WithContext(ctx).Session(&gorm.Session{AllowGlobalUpdate: true}).
		Unscoped().Delete(&entities.ToDo{}).Error; err != nil {
		return fmt.Errorf("failed to delete todos: %w", err)
	}

	if err := db.WithContext(ctx).Session(&gorm.Session{AllowGlobalUpdate: true}).
		Unscoped().Delete(&entities.ToDoList{}).Error; err != nil {
		return fmt.Errorf("failed to delete todo lists: %w", err)
	}

	if err := db.WithContext(ctx).Session(&gorm.Session{AllowGlobalUpdate: true}).
		Unscoped().Delete(&entities.Project{}).Error; err != nil {
		return fmt.Errorf("failed to delete projects: %w", err)
	}

	if err := db.WithContext(ctx).Session(&gorm.Session{AllowGlobalUpdate: true}).
		Unscoped().Delete(&entities.Settings{}).Error; err != nil {
		return fmt.Errorf("failed to delete settings: %w", err)
	}

	log.Println("✓ Database cleaned")
	return nil
}

func seedSettings(tx *gorm.DB, ctx context.Context) error {
	var count int64
	tx.WithContext(ctx).Model(&entities.Settings{}).Count(&count)
	if count > 0 {
		log.Println("⚙️  Settings already exist, skipping")
		return nil
	}

	defaultSettings := entities.Settings{
		ActiveThemeID: "dark",
	}

	if err := tx.WithContext(ctx).Create(&defaultSettings).Error; err != nil {
		return fmt.Errorf("failed to create settings: %w", err)
	}

	log.Printf("✅ Created default settings (Theme: %s)", defaultSettings.ActiveThemeID)
	return nil
}

func seedProjects(tx *gorm.DB, ctx context.Context) error {
	var count int64
	tx.WithContext(ctx).Model(&entities.Project{}).Count(&count)
	if count > 0 {
		log.Println("📁 Projects already exist, skipping")
		return nil
	}

	projects := []*entities.Project{
		{Name: "Work", Color: "#FF6B6B"},
		{Name: "Personal", Color: "#4ECDC4"},
		{Name: "Homelab", Color: "#45B7D1"},
	}

	for _, p := range projects {
		if err := tx.WithContext(ctx).Create(p).Error; err != nil {
			return fmt.Errorf("failed to create project %s: %w", p.Name, err)
		}
		log.Printf("✅ Created project: %s (ID: %d)", p.Name, p.ID)
	}

	return nil
}

func seedToDoLists(tx *gorm.DB, ctx context.Context) error {
	var count int64
	tx.WithContext(ctx).Model(&entities.ToDoList{}).Count(&count)
	if count > 0 {
		log.Println("📋 ToDo lists already exist, skipping")
		return nil
	}

	todoLists := []*entities.ToDoList{
		{Name: "Urgent", Color: "#FF4757"},
		{Name: "Daily Tasks", Color: "#FFA502"},
		{Name: "Backlog", Color: "#2ED573"},
		{Name: "Homelab Setup", Color: "#1E90FF"},
	}

	for _, tl := range todoLists {
		if err := tx.WithContext(ctx).Create(tl).Error; err != nil {
			return fmt.Errorf("failed to create todolist %s: %w", tl.Name, err)
		}
		log.Printf("✅ Created todolist: %s (ID: %d)", tl.Name, tl.ID)
	}

	return nil
}

func seedToDos(tx *gorm.DB, ctx context.Context) error {
	var count int64
	tx.WithContext(ctx).Model(&entities.ToDo{}).Count(&count)
	if count > 0 {
		log.Println("✓ ToDos already exist, skipping")
		return nil
	}

	var projects []entities.Project
	if err := tx.WithContext(ctx).Order("id").Find(&projects).Error; err != nil {
		return fmt.Errorf("failed to load projects: %w", err)
	}

	if len(projects) < 3 {
		return fmt.Errorf("not enough projects to seed todos (need 3, have %d)", len(projects))
	}

	var todoLists []entities.ToDoList
	if err := tx.WithContext(ctx).Order("id").Find(&todoLists).Error; err != nil {
		return fmt.Errorf("failed to load todo lists: %w", err)
	}

	if len(todoLists) < 4 {
		return fmt.Errorf("not enough todo lists to seed todos (need 4, have %d)", len(todoLists))
	}

	todos := []*entities.ToDo{
		{
			ProjectID:   projects[0].ID,
			ToDoListID:  todoLists[0].ID,
			Name:        "Deploy microservice v2",
			Description: strPtr("Update Docker → Alpine 3.20 + Ansible deploy"),
			Priority:    enums.High,
			Status:      enums.Pending,
			Color:       "#FF4757",
			Done:        false,
			DueDate:     tomorrow(),
		},
		{
			ProjectID:  projects[0].ID,
			ToDoListID: todoLists[1].ID,
			Name:       "Code review PR #456",
			Priority:   enums.Medium,
			Status:     enums.InProgress,
			Color:      "#FFA502",
			Done:       false,
		},
		{
			ProjectID:  projects[1].ID,
			ToDoListID: todoLists[2].ID,
			Name:       "Grocery shopping",
			Priority:   enums.Low,
			Status:     enums.Done,
			Color:      "#2ED573",
			Done:       true,
		},
		{
			ProjectID:   projects[2].ID,
			ToDoListID:  todoLists[3].ID,
			Name:        "Proxmox NFS backup",
			Description: strPtr("NFSv4 + MergerFS + daily cron"),
			Priority:    enums.High,
			Status:      enums.Pending,
			Color:       "#1E90FF",
			Done:        false,
			DueDate:     weekFromNow(),
		},
		{
			ProjectID: projects[2].ID,
			Name:      "Update Ansible Galaxy roles",
			Priority:  enums.Medium,
			Status:    enums.Pending,
			Color:     "#74B9FF",
			Done:      false,
		},
	}

	for _, todo := range todos {
		if err := tx.WithContext(ctx).Create(todo).Error; err != nil {
			return fmt.Errorf("failed to create todo '%s': %w", todo.Name, err)
		}

		projectName := ""
		for _, p := range projects {
			if p.ID == todo.ProjectID {
				projectName = p.Name
				break
			}
		}
		log.Printf("✅ Created todo: %s (ID: %d, Project: %s)", todo.Name, todo.ID, projectName)
	}

	return nil
}

func verifySeeding(db *gorm.DB, ctx context.Context) error {
	log.Println("\n=== VERIFICATION ===")

	var settingsCount int64
	db.WithContext(ctx).Model(&entities.Settings{}).Count(&settingsCount)
	log.Printf("⚙️  Settings: %d", settingsCount)

	var projectCount int64
	db.WithContext(ctx).Model(&entities.Project{}).Count(&projectCount)
	log.Printf("📁 Projects: %d", projectCount)

	allLists, err := gorm.G[entities.ToDoList](db).Find(ctx)
	if err != nil {
		return err
	}
	log.Printf("📋 Total lists: %d", len(allLists))

	var workProject entities.Project
	if err := db.WithContext(ctx).Where("name = ?", "Work").First(&workProject).Error; err == nil {
		workTodos, err := gorm.G[entities.ToDo](db).
			Where(generated.ToDo.ProjectID.Eq(workProject.ID)).
			Preload("Project", nil).
			Preload("ToDoList", nil).
			Find(ctx)
		if err == nil {
			log.Printf("💼 Work project: %d todos", len(workTodos))
		}
	}

	urgent, err := gorm.G[entities.ToDo](db).
		Where(generated.ToDo.Priority.WithName(enums.High.String())).
		Where(generated.ToDo.Status.WithName(enums.Pending.String())).
		Find(ctx)
	if err == nil {
		log.Printf("🔥 High priority pending: %d todos", len(urgent))
	}

	done, err := gorm.G[entities.ToDo](db).
		Where(generated.ToDo.Done.Eq(true)).
		Find(ctx)
	if err == nil {
		log.Printf("✅ Completed: %d todos", len(done))
	}

	log.Println("===================")
	return nil
}

func ResetAndSeed(dbService *DbService) error {
	if err := CleanDatabase(dbService); err != nil {
		return err
	}
	return Seed(dbService)
}

// Utility functions
func strPtr(s string) *string {
	return &s
}

func tomorrow() *time.Time {
	t := time.Now().Truncate(time.Hour * 24).Add(24 * time.Hour)
	return &t
}

func weekFromNow() *time.Time {
	t := time.Now().Truncate(time.Hour * 24).Add(7 * 24 * time.Hour)
	return &t
}
