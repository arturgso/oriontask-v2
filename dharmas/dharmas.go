package dharmas

import (
	"time"

	"github.com/google/uuid"
)

type Dharmas struct {
	ID        uuid.UUID
	Name      string
	createdAt time.Time
}
