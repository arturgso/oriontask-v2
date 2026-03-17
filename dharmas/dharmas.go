package dharmas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dharmas struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(60);not null" json:"name"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (d *Dharmas) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return
}
