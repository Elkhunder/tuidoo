package services

import (
	"context"
	"log"
	"sync"
	"time"
	e "tuidoo/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	once    sync.Once
	initErr error
)

type DbService struct {
	db *gorm.DB
}

func (d *DbService) Connect() error {
	once.Do(func() {
		db, initErr = gorm.Open(sqlite.Open("test_tuidoo.db"), &gorm.Config{})
		if initErr != nil {
			log.Printf("Failed to connect to database: %v", initErr)
			return
		}

		initErr = db.AutoMigrate(
			&e.Settings{},
			&e.Project{},
			&e.ToDoList{},
			&e.ToDo{},
		)
		if initErr != nil {
			log.Printf("Failed to run migrations: %v", initErr)
			return
		}

		log.Println("Database connected and migrated successfully")
	})
	if initErr != nil {
		return initErr
	}

	d.db = db
	return nil
}

// Close closes the database connection
func (d *DbService) Close() error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

// GetDB returns the database instance
func (d *DbService) GetDB() *gorm.DB {
	return d.db
}

// NewContext creates a context with timeout for database operations
func (d *DbService) NewContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

var Db = &DbService{}
