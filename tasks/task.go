package tasks

import (
	"oriontask-v2/dharmas"
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	Todo       TaskStatus = "TODO"
	InProgress TaskStatus = "IN_PROGRESS"
	Blocked    TaskStatus = "BLOCKED"
	Completed  TaskStatus = "COMPLETED"
	Cancelled  TaskStatus = "CANCELLED"
)

type Task struct {
	ID          uuid.UUID        `gorm:"primaryKey" json:"id"`
	DharmaID    uuid.UUID        `gorm:"not null" json:"dharmaId"`
	Dharma      *dharmas.Dharmas `gorm:"foreignKey:DharmaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Title       string           `gorm:"type:varchar(100);not null" json:"title"`
	Description string           `gorm:"type:text;" json:"description"`
	Status      TaskStatus       `gorm:"type:varchar(20);not null;default:'TODO'" json:"status"`
	CompletedAt *time.Time       `gorm:"type:timestamp" json:"completedAt"`
	CreatedAt   time.Time        `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time        `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ts TaskStatus) IsValid() bool {
	switch ts {
	case Todo, InProgress, Blocked, Completed, Cancelled:
		return true
	}
	return false
}
