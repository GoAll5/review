package domain

import (
	"github.com/google/uuid"
	"time"
)

type Review struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	UserID    uuid.UUID
	Rate      int
	Text      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
