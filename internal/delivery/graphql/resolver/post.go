package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input domain.CreatePostInput) (string, error) {
	input.AuthorID = ctx.Value(domain.UserCtx).(string)

	return r.service.Post.CreatePost(ctx, input)
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	return r.service.Post.DeletePost(ctx, id)
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input domain.UpdatePostInput) (bool, error) {
	return r.service.Post.UpdatePost(ctx, input)
}

func (r *queryResolver) Post(ctx context.Context, id string) (*domain.Post, error) {
	return r.service.Post.GetPost(ctx, id)
}
