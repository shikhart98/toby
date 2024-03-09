package repo_user

import (
	domain_user "toby/src/internal/domain/user"
)

func (u *UserRepository) toDomain() *domain_user.UserWithID {
	return &domain_user.UserWithID{
		ID: u.ID,
		User: domain_user.User{
			Name:  u.Name,
			Email: u.Email,
		},
	}
}

func NewRepository(user *domain_user.CreateUserPayload) UserRepository {
	return UserRepository{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
