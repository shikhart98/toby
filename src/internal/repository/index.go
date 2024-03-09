package repository

import (
	"github.com/google/uuid"
	"time"
)

type BaseRepository struct {
	ID        uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt *time.Time `gorm:"not null;autoCreateTime:nano"`
	UpdatedAt *time.Time `gorm:"not null;autoUpdateTime:nano"`
}
