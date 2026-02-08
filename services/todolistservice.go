package services

import (
	"fmt"
	"tuidoo/entities"
)

type ToDoListService struct {
	db *DbService
}

func NewToDoListService(dbService *DbService) *ToDoListService {
	return &ToDoListService{db: dbService}
}

func (tls *ToDoListService) Create(list *entities.ToDoList) error {
	ctx, cancel := tls.db.NewContext()
	defer cancel()

	if list.Name == "" {
		return fmt.Errorf("list name cannot be empty")
	}

	if err := tls.db.GetDB().WithContext(ctx).Create(list).Error; err != nil {
		return fmt.Errorf("failed to create list: %w", err)
	}

	return nil
}

func (tls *ToDoListService) GetByID(id uint, includeToDos bool) (*entities.ToDoList, error) {
	ctx, cancel := tls.db.NewContext()
	defer cancel()

	query := tls.db.GetDB().WithContext(ctx)

	if includeToDos {
		query = query.Preload("ToDos")
	}

	var list entities.ToDoList
	if err := query.First(&list, id).Error; err != nil {
		return nil, fmt.Errorf("failed to get list: %w", err)
	}

	return &list, nil
}

func (tls *ToDoListService) GetAll(includeToDos bool) ([]entities.ToDoList, error) {
	ctx, cancel := tls.db.NewContext()
	defer cancel()

	query := tls.db.GetDB().WithContext(ctx)

	if includeToDos {
		query = query.Preload("ToDos")
	}

	var lists []entities.ToDoList
	if err := query.Find(&lists).Error; err != nil {
		return nil, fmt.Errorf("failed to get lists: %w", err)
	}

	return lists, nil
}

func (tls *ToDoListService) Update(list *entities.ToDoList) error {
	ctx, cancel := tls.db.NewContext()
	defer cancel()

	if list.Name == "" {
		return fmt.Errorf("list name cannot be empty")
	}

	if err := tls.db.GetDB().WithContext(ctx).Save(list).Error; err != nil {
		return fmt.Errorf("failed to update list: %w", err)
	}

	return nil
}

func (tls *ToDoListService) Delete(id uint) error {
	ctx, cancel := tls.db.NewContext()
	defer cancel()

	result := tls.db.GetDB().WithContext(ctx).Delete(&entities.ToDoList{}, id)

	if result.Error != nil {
		return fmt.Errorf("failed to delete list: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("list with ID %d not found", id)
	}

	return nil
}

func (tls *ToDoListService) Search(query string) ([]entities.ToDoList, error) {
	ctx, cancel := tls.db.NewContext()
	defer cancel()

	var lists []entities.ToDoList
	if err := tls.db.GetDB().WithContext(ctx).
		Where("name LIKE ?", "%"+query+"%").
		Find(&lists).Error; err != nil {
		return nil, fmt.Errorf("failed to search lists: %w", err)
	}

	return lists, nil
}
