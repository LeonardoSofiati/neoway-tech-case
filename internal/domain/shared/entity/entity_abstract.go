package shared

import (
	"time"

	"github.com/rs/xid"
)

type BaseEntity struct {
	ID        string    `json:"id" gorm:"primaryKey;size:50;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
}

func NewBaseEntity() BaseEntity {
	return BaseEntity{
		ID:        xid.New().String(),
		CreatedAt: time.Now(),
	}
}
