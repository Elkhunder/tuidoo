package services

import (
	"fmt"
	"tuidoo/entities"
)

type ProjectService struct {
	db *DbService
}

func NewProjectService(dbService *DbService) *ProjectService {
	return &ProjectService{db: dbService}
}

func (ps *ProjectService) Create(project *entities.Project) error {
	ctx, cancel := ps.db.NewContext()
	defer cancel()

	if project.Name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	if err := ps.db.GetDB().WithContext(ctx).Create(project).Error; err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}

	return nil
}

func (ps *ProjectService) GetByID(id uint, includeToDos bool) (*entities.Project, error) {
	ctx, cancel := ps.db.NewContext()
	defer cancel()

	query := ps.db.GetDB().WithContext(ctx)

	if includeToDos {
		query = query.Preload("ToDos")
	}

	var project entities.Project
	if err := query.First(&project, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}

	return &project, nil
}

func (ps *ProjectService) GetAll(includeToDos bool) ([]entities.Project, error) {
	ctx, cancel := ps.db.NewContext()
	defer cancel()

	query := ps.db.GetDB().WithContext(ctx)

	if includeToDos {
		query = query.Preload("ToDos")
	}

	var projects []entities.Project
	if err := query.Find(&projects).Error; err != nil {
		return nil, fmt.Errorf("failed to get projects: %w", err)
	}

	return projects, nil
}

func (ps *ProjectService) Update(project *entities.Project) error {
	ctx, cancel := ps.db.NewContext()
	defer cancel()

	if project.Name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	if err := ps.db.GetDB().WithContext(ctx).Save(project).Error; err != nil {
		return fmt.Errorf("failed to update project: %w", err)
	}

	return nil
}

func (ps *ProjectService) Delete(id uint) error {
	ctx, cancel := ps.db.NewContext()
	defer cancel()

	result := ps.db.GetDB().WithContext(ctx).Delete(&entities.Project{}, id)

	if result.Error != nil {
		return fmt.Errorf("failed to delete project: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("project with ID %d not found", id)
	}

	return nil
}

func (ps *ProjectService) Search(query string) ([]entities.Project, error) {
	ctx, cancel := ps.db.NewContext()
	defer cancel()

	var projects []entities.Project
	if err := ps.db.GetDB().WithContext(ctx).
		Where("name LIKE ?", "%"+query+"%").
		Find(&projects).Error; err != nil {
		return nil, fmt.Errorf("failed to search projects: %w", err)
	}

	return projects, nil
}
