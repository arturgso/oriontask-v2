package dharmas

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

// MockDharmaRepository is a manual mock of the DharmaRepository interface
type MockDharmaRepository struct {
	CreateFunc  func(ctx context.Context, dharma *Dharmas) error
	GetByIDFunc func(ctx context.Context, id uuid.UUID) (*Dharmas, error)
	ListAllFunc func(ctx context.Context) ([]Dharmas, error)
	DeleteFunc  func(ctx context.Context, id uuid.UUID) error
}

func (m *MockDharmaRepository) Create(ctx context.Context, dharma *Dharmas) error {
	return m.CreateFunc(ctx, dharma)
}
func (m *MockDharmaRepository) GetByID(ctx context.Context, id uuid.UUID) (*Dharmas, error) {
	return m.GetByIDFunc(ctx, id)
}
func (m *MockDharmaRepository) ListAll(ctx context.Context) ([]Dharmas, error) {
	return m.ListAllFunc(ctx)
}
func (m *MockDharmaRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return m.DeleteFunc(ctx, id)
}

func TestDharmaAppService_CreateDharma(t *testing.T) {
	tests := []struct {
		name          string
		inputName     string
		mockReturnErr error
		wantErr       bool
		errMsg        string
	}{
		{
			name:      "Success - Valid Name",
			inputName: "Work",
			wantErr:   false,
		},
		{
			name:      "Error - Empty Name",
			inputName: "   ",
			wantErr:   true,
			errMsg:    "O nome do Dharma não pode estar vazaio",
		},
		{
			name:      "Error - Name Too Long",
			inputName: "This name is definitely longer than sixty characters to trigger the validation error",
			wantErr:   true,
			errMsg:    "O nome do Dharma não pode ter mais de 60 caracteres",
		},
		{
			name:          "Error - Repository Failure",
			inputName:     "Health",
			mockReturnErr: errors.New("db error"),
			wantErr:       true,
			errMsg:        "db error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockDharmaRepository{
				CreateFunc: func(ctx context.Context, dharma *Dharmas) error {
					return tt.mockReturnErr
				},
			}

			service := NewDharmaService(mockRepo)
			result, err := service.CreateDharma(tt.inputName)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDharma() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("CreateDharma() error message = %v, want %v", err.Error(), tt.errMsg)
			}

			if !tt.wantErr && result.Name != tt.inputName {
				t.Errorf("CreateDharma() result name = %v, want %v", result.Name, tt.inputName)
			}
		})
	}
}

func TestDharmaAppService_DeleteDharma(t *testing.T) {
	testID := uuid.New()

	tests := []struct {
		name          string
		id            uuid.UUID
		mockGetResult *Dharmas
		mockGetErr    error
		mockDelErr    error
		wantErr       bool
		errMsg        string
	}{
		{
			name:          "Success",
			id:            testID,
			mockGetResult: &Dharmas{ID: testID},
			wantErr:       false,
		},
		{
			name:    "Error - Nil ID",
			id:      uuid.Nil,
			wantErr: true,
			errMsg:  "ID inválido fornecido",
		},
		{
			name:       "Error - Not Found",
			id:         testID,
			mockGetErr: errors.New("not found"),
			wantErr:    true,
			errMsg:     "não é possível deletar: dharma não existe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockDharmaRepository{
				GetByIDFunc: func(ctx context.Context, id uuid.UUID) (*Dharmas, error) {
					return tt.mockGetResult, tt.mockGetErr
				},
				DeleteFunc: func(ctx context.Context, id uuid.UUID) error {
					return tt.mockDelErr
				},
			}

			service := NewDharmaService(mockRepo)
			err := service.DeleteDharma(tt.id)

			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteDharma() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("DeleteDharma() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}
