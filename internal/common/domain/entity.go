package common_domain

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
