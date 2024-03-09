package domain_feedback

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Feedback struct {
	Provider    uuid.UUID `json:"provider"`
	Receiver    uuid.UUID `json:"receiver"`
	CreatedAt   time.Time `json:"created_at"`
	Data        string    `json:"data"`
	isAnonymous bool      `json:"is_anonymous"`
}

type FeedbackWithID struct {
	ID       uuid.UUID `json:"id"`
	Feedback `json:",inline"`
}

type Usecase interface {
	Publish(ctx context.Context, feedback *Feedback) (*FeedbackWithID, error)
	Remove(ctx context.Context, feedback *FeedbackWithID) error
	GetUserFeedback(ctx context.Context, userId uuid.UUID) ([]*FeedbackWithID, error)
	GetPublishedFeedbackList(ctx context.Context, userId uuid.UUID) ([]*FeedbackWithID, error)
}

type Repository interface {
	Create(ctx context.Context, feedback *Feedback) (*FeedbackWithID, error)
	SoftDelete(ctx context.Context, id uuid.UUID) error
	ListByProvider(ctx context.Context, providerId uuid.UUID) ([]*FeedbackWithID, error)
	ListByReceiver(ctx context.Context, receiverId uuid.UUID) ([]*FeedbackWithID, error)
}
