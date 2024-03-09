package repo_user

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	domain_user "toby/src/internal/domain/user"
	"toby/src/internal/repository"
)

type UserRepository struct {
	Name     string `gorm:"not null;"`
	Email    string `gorm:"not null;"`
	Password string `gorm:"not null;"`

	repository.BaseRepository
}

func (UserRepository) TableName() string {
	return "user"
}

type UserRepositoryImplementation struct {
	db *gorm.DB
}

func (u *UserRepositoryImplementation) Create(ctx context.Context, user *domain_user.CreateUserPayload) (*domain_user.UserWithID, error) {
	userRepoObj := NewRepository(user)
	err := u.db.WithContext(ctx).Create(&userRepoObj).Error
	if err != nil {
		return nil, err
	}
	return userRepoObj.toDomain(), nil
}

func (u *UserRepositoryImplementation) Update(ctx context.Context, user *domain_user.UserWithID) (*domain_user.UserWithID, error) {

}

func (u *UserRepositoryImplementation) Get(ctx context.Context, userId uuid.UUID) (*domain_user.UserWithID, error) {

}
