package tasks

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type TaskRepository interface {
	Save(ctx context.Context, task *Task) error
	Update(ctx context.Context, task *Task) error
	// TODO: FindByProjectID(ctx context.Context, projectID uuid.UUID) ([]Task, error)
	// TODO: FindByMilestoneID(ctx context.Context, milestoneID uuid.UUID) ([]Task, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Task, error)
	FindByDharmaID(ctx context.Context, dharmaID uuid.UUID) ([]Task, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(task Task) (*Task, error) {
	title := strings.TrimSpace(task.Title)

	if title == "" {
		return nil, errors.New("O título da task não deve estar vazio")
	}

	newTask := &Task{
		ID:          uuid.New(),
		DharmaID:    task.DharmaID,
		Title:       strings.TrimSpace(task.Title),
		Description: strings.TrimSpace(task.Description),
		Status:      task.Status,
	}

	err := s.repo.Save(context.Background(), newTask)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (s *TaskService) FindTaskByID(id uuid.UUID) (*Task, error) {
	if id == uuid.Nil {
		return nil, errors.New("ID inválido fornecido")
	}

	return s.repo.FindByID(context.Background(), id)
}

func (s *TaskService) FindTasksByDharmaID(id uuid.UUID) ([]Task, error) {
	return s.repo.FindByDharmaID(context.Background(), id)
}

// TODO: func (s *TaskService) FindTasksByProjectID(projectID uuid.UUID) ([]Task, error) { ... }
// TODO: func (s *TaskService) FindTasksByMilestoneID(milestoneID uuid.UUID) ([]Task, error) { ... }

func (s *TaskService) CompleteTask(id uuid.UUID) (*Task, error) {
	if id == uuid.Nil {
		return nil, errors.New("ID inválido fornecido")
	}

	task, err := s.repo.FindByID(context.Background(), id)
	if err != nil {
		return nil, errors.New("Task não encontrada")
	}

	if task.CompletedAt != nil {
		return nil, errors.New("Task já marcada como concluída")
	}

	now := time.Now()
	task.CompletedAt = &now
	task.Status = "COMPLETED"

	if err := s.repo.Update(context.Background(), task); err != nil {
		return nil, err
	}

	return task, nil
}

// TODO: func (s *TaskService) PostponeTask(id uuid.UUID) (*Task, error) { ... }

func (s *TaskService) DeleteTask(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("ID inválido fornecido")
	}

	_, err := s.repo.FindByID(context.Background(), id)
	if err != nil {
		return errors.New("Não foi possível deltear: task não existe")
	}

	return s.repo.Delete(context.Background(), id)
}
