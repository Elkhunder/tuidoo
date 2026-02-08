package services

import (
	"fmt"
	"time"
	"tuidoo/entities"
	"tuidoo/enums"
	"tuidoo/generated"

	"gorm.io/gorm"
)

type ToDoService struct {
	db *DbService
}

func NewToDoService(dbService *DbService) *ToDoService {
	return &ToDoService{db: dbService}
}

// Create creates a new todo
func (ts *ToDoService) Create(todo *entities.ToDo) error {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	if err := ts.validate(todo); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := ts.db.GetDB().WithContext(ctx).Create(todo).Error; err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
	}

	return nil
}

// GetByID retrieves a todo by ID with optional preloading
func (ts *ToDoService) GetByID(id uint, preload bool) (*entities.ToDo, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	query := ts.db.GetDB().WithContext(ctx)

	if preload {
		query = query.Preload("Project").Preload("ToDoList")
	}

	var todo entities.ToDo
	if err := query.First(&todo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("todo with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to get todo: %w", err)
	}

	return &todo, nil
}

// GetAll retrieves all todos with optional filters and preloading
func (ts *ToDoService) GetAll(preload bool) ([]entities.ToDo, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	query := ts.db.GetDB().WithContext(ctx)

	if preload {
		query = query.Preload("Project").Preload("ToDoList")
	}

	var todos []entities.ToDo
	if err := query.Find(&todos).Error; err != nil {
		return nil, fmt.Errorf("failed to get todos: %w", err)
	}

	return todos, nil
}

// GetByProject retrieves all todos for a specific project
func (ts *ToDoService) GetByProject(projectID uint, preload bool) ([]entities.ToDo, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	todos, err := gorm.G[entities.ToDo](ts.db.GetDB()).
		Where(generated.ToDo.ProjectID.Eq(projectID)).
		Preload("Project", nil).
		Preload("ToDoList", nil).
		Find(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get todos by project: %w", err)
	}

	return todos, nil
}

// GetByList retrieves all todos for a specific list
func (ts *ToDoService) GetByList(listID uint, preload bool) ([]entities.ToDo, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	todos, err := gorm.G[entities.ToDo](ts.db.GetDB()).
		Where(generated.ToDo.ToDoListID.Eq(listID)).
		Preload("Project", nil).
		Preload("ToDoList", nil).
		Find(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get todos by list: %w", err)
	}

	return todos, nil
}

// GetByStatus retrieves todos by status
func (ts *ToDoService) GetByStatus(status enums.Status, preload bool) ([]entities.ToDo, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	todos, err := gorm.G[entities.ToDo](ts.db.GetDB()).
		Where(generated.ToDo.Status.WithName(status.String())).
		Preload("Project", nil).
		Preload("ToDoList", nil).
		Find(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get todos by status: %w", err)
	}

	return todos, nil
}

// GetByPriority retrieves todos by priority
func (ts *ToDoService) GetByPriority(priority enums.Priority, preload bool) ([]entities.ToDo, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	todos, err := gorm.G[entities.ToDo](ts.db.GetDB()).
		Where(generated.ToDo.Priority.WithName(priority.String())).
		Preload("Project", nil).
		Preload("ToDoList", nil).
		Find(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get todos by priority: %w", err)
	}

	return todos, nil
}

// GetCompleted retrieves all completed todos
func (ts *ToDoService) GetCompleted(preload bool) ([]entities.ToDo, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	todos, err := gorm.G[entities.ToDo](ts.db.GetDB()).
		Where(generated.ToDo.Done.Eq(true)).
		Preload("Project", nil).
		Preload("ToDoList", nil).
		Find(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get completed todos: %w", err)
	}

	return todos, nil
}

// GetPending retrieves all pending (not done) todos
func (ts *ToDoService) GetPending(preload bool) ([]entities.ToDo, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	todos, err := gorm.G[entities.ToDo](ts.db.GetDB()).
		Where(generated.ToDo.Done.Eq(false)).
		Preload("Project", nil).
		Preload("ToDoList", nil).
		Find(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get pending todos: %w", err)
	}

	return todos, nil
}

// GetOverdue retrieves todos past their due date that aren't completed
func (ts *ToDoService) GetOverdue(preload bool) ([]entities.ToDo, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	now := time.Now()

	query := ts.db.GetDB().WithContext(ctx).
		Where("due_date < ? AND done = ?", now, false)

	if preload {
		query = query.Preload("Project").Preload("ToDoList")
	}

	var todos []entities.ToDo
	if err := query.Find(&todos).Error; err != nil {
		return nil, fmt.Errorf("failed to get overdue todos: %w", err)
	}

	return todos, nil
}

// Update updates an existing todo
func (ts *ToDoService) Update(todo *entities.ToDo) error {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	if err := ts.validate(todo); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := ts.db.GetDB().WithContext(ctx).Save(todo).Error; err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}

	return nil
}

// UpdateStatus updates just the status of a todo
func (ts *ToDoService) UpdateStatus(id uint, status enums.Status) error {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	result := ts.db.GetDB().WithContext(ctx).
		Model(&entities.ToDo{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return fmt.Errorf("failed to update status: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	return nil
}

// MarkAsComplete marks a todo as completed
func (ts *ToDoService) MarkAsComplete(id uint) error {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	result := ts.db.GetDB().WithContext(ctx).
		Model(&entities.ToDo{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"done":   true,
			"status": enums.Done,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to mark todo as complete: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	return nil
}

// MarkAsIncomplete marks a todo as not completed
func (ts *ToDoService) MarkAsIncomplete(id uint) error {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	result := ts.db.GetDB().WithContext(ctx).
		Model(&entities.ToDo{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"done":   false,
			"status": enums.Pending,
		})

	if result.Error != nil {
		return fmt.Errorf("failed to mark todo as incomplete: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	return nil
}

// Delete soft deletes a todo
func (ts *ToDoService) Delete(id uint) error {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	result := ts.db.GetDB().WithContext(ctx).Delete(&entities.ToDo{}, id)

	if result.Error != nil {
		return fmt.Errorf("failed to delete todo: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	return nil
}

// HardDelete permanently deletes a todo
func (ts *ToDoService) HardDelete(id uint) error {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	result := ts.db.GetDB().WithContext(ctx).Unscoped().Delete(&entities.ToDo{}, id)

	if result.Error != nil {
		return fmt.Errorf("failed to hard delete todo: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("todo with ID %d not found", id)
	}

	return nil
}

// Search searches todos by name (case-insensitive)
func (ts *ToDoService) Search(query string, preload bool) ([]entities.ToDo, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	dbQuery := ts.db.GetDB().WithContext(ctx).
		Where("name LIKE ?", "%"+query+"%")

	if preload {
		dbQuery = dbQuery.Preload("Project").Preload("ToDoList")
	}

	var todos []entities.ToDo
	if err := dbQuery.Find(&todos).Error; err != nil {
		return nil, fmt.Errorf("failed to search todos: %w", err)
	}

	return todos, nil
}

// Count returns the total number of todos
func (ts *ToDoService) Count() (int64, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	var count int64
	if err := ts.db.GetDB().WithContext(ctx).Model(&entities.ToDo{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count todos: %w", err)
	}

	return count, nil
}

// CountByStatus returns the count of todos by status
func (ts *ToDoService) CountByStatus(status enums.Status) (int64, error) {
	ctx, cancel := ts.db.NewContext()
	defer cancel()

	var count int64
	if err := ts.db.GetDB().WithContext(ctx).
		Model(&entities.ToDo{}).
		Where("status = ?", status).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count todos by status: %w", err)
	}

	return count, nil
}

// validate performs basic validation on a todo
func (ts *ToDoService) validate(todo *entities.ToDo) error {
	if todo.Name == "" {
		return fmt.Errorf("todo name cannot be empty")
	}

	if todo.ProjectID == 0 {
		return fmt.Errorf("todo must be assigned to a project")
	}

	return nil
}
