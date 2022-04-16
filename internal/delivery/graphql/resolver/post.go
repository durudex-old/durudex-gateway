package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input domain.CreatePostInput) (string, error) {
	return "", nil
}

func (r *queryResolver) GetPost(ctx context.Context, id string) (*domain.Post, error) {
	return &domain.Post{}, nil
}
