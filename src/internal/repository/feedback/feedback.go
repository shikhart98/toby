package repo_feedback

import (
	"github.com/google/uuid"
	"toby/src/internal/repository"
	repo_user "toby/src/internal/repository/user"
)

type FeedbackRepository struct {
	ProviderID  uuid.UUID `gorm:"not null;"`
	ReceiverID  uuid.UUID `gorm:"not null;"`
	Data        string    `gorm:"not null;"`
	IsAnonymous bool      `gorm:"not null;default:false"`

	Provider repo_user.UserRepository `gorm:"foreignKey:ProviderID;constraint:OnDelete:CASCADE"`
	Receiver repo_user.UserRepository `gorm:"foreignKey:ReceiverID;constraint:OnDelete:CASCADE"`
	repository.BaseRepository
}

func (FeedbackRepository) TableName() string {
	return "feedback"
}
