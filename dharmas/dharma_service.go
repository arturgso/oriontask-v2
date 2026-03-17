package dharmas

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
)

type DharmaAppService struct {
	repo DharmaRepository
}

func NewDharmaService(repo DharmaRepository) *DharmaAppService {
	return &DharmaAppService{
		repo: repo,
	}
}

func (s *DharmaAppService) CreateDharma(name string) (*Dharmas, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("O nome do Dharma não pode estar vazio")
	}

	if len(name) > 60 {
		return nil, errors.New("O nome do Dharma não pode ter mais de 60 caracteres")
	}

	dharma := &Dharmas{
		ID:   uuid.New(),
		Name: name,
	}

	err := s.repo.Create(context.Background(), dharma)
	if err != nil {
		return nil, err
	}

	return dharma, nil
}

func (s *DharmaAppService) GetDharmaByID(id uuid.UUID) (*Dharmas, error) {
	if id == uuid.Nil {
		return nil, errors.New("ID inválido fornecido")
	}

	return s.repo.GetByID(context.Background(), id)
}

func (s *DharmaAppService) ListAllDharmas() ([]Dharmas, error) {
	return s.repo.ListAll(context.Background())
}

func (s *DharmaAppService) DeleteDharma(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("ID inválido fornecido")
	}

	_, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return errors.New("não é possível deletar: dharma não existe")
	}

	return s.repo.Delete(context.Background(), id)
}
