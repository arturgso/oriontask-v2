package tasks

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Save(ctx context.Context, task *Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *taskRepository) FindByID(ctx context.Context, id uuid.UUID) (*Task, error) {
	var task Task

	err := r.db.WithContext(ctx).First(&task, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task não encontrada")
		}
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) ListAll(ctx context.Context) ([]Task, error) {
	var task []Task
	err := r.db.WithContext(ctx).Find(&task).Error
	return task, err
}

func (r *taskRepository) FindByDharmaID(ctx context.Context, dharmaID uuid.UUID) ([]Task, error) {
	var tasks []Task
	err := r.db.WithContext(ctx).Where("dharma_id = ?", dharmaID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&Task{}, "id = ?", id).Error
}
