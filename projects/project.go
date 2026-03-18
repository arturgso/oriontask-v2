package projects

import (
	"oriontask-v2/dharmas"
	"time"

	"github.com/google/uuid"
)

type ProjectStatus string

const (
	Planning  ProjectStatus = "PLANNING"
	Active    ProjectStatus = "ACTIVE"
	OnHold    ProjectStatus = "ON_HOLD"
	Completed ProjectStatus = "COMPLETED"
)

type Project struct {
	ID          uuid.UUID        `gorm:"primaryKey" json:"id"`
	DharmaID    uuid.UUID        `gorm:"not null" json:"dharmaId"`
	Dharma      *dharmas.Dharmas `gorm:"foreignKey:DharmaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Title       string           `gorm:"type:varchar(40);not null" json:"title"`
	Description string           `gorm:"type:text;" json:"description"`
	Status      ProjectStatus    `gorm:"type:text;not null;default:'PLANNING'" json:"status"`
	CreatedAt   time.Time        `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time        `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
