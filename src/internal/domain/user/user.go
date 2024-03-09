package domain_user

import (
	"context"
	"github.com/google/uuid"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserPayload struct {
	User     `json:",inline"`
	Password string `json:"password"`
}

type UpdateUserPayload struct {
	Name string `json:"name"`
}

type UserWithID struct {
	ID   uuid.UUID `json:"id"`
	User `json:",inline"`
}

type Usecase interface {
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
}

type Repository interface {
	Create(ctx context.Context, user *User) (*UserWithID, error)
	Update(ctx context.Context, user *UserWithID) (*UserWithID, error)
	Get(ctx context.Context, userId uuid.UUID) (*UserWithID, error)
}
