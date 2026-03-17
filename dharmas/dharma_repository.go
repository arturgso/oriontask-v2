package dharmas

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DharmaRepository interface {
	Create(ctx context.Context, dharma *Dharmas) error
	GetByID(ctx context.Context, id uuid.UUID) (*Dharmas, error)
	ListAll(ctx context.Context) ([]Dharmas, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type dharmaRepository struct {
	db *gorm.DB
}

func NewDharmaRepository(db *gorm.DB) DharmaRepository {
	return &dharmaRepository{
		db: db,
	}
}

func (r *dharmaRepository) Create(ctx context.Context, dharma *Dharmas) error {
	return r.db.WithContext(ctx).Create(dharma).Error
}

func (r *dharmaRepository) GetByID(ctx context.Context, id uuid.UUID) (*Dharmas, error) {
	var dharma Dharmas

	err := r.db.WithContext(ctx).First(&dharma, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("dharma não encontrado")
		}
		return nil, err
	}
	return &dharma, nil
}

func (r *dharmaRepository) ListAll(ctx context.Context) ([]Dharmas, error) {
	var dharmas []Dharmas
	err := r.db.WithContext(ctx).Find(&dharmas).Error
	return dharmas, err
}

func (r *dharmaRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&Dharmas{}, "id = ?", id).Error
}
