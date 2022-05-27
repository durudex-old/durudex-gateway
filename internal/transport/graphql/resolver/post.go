package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
	"github.com/durudex/durudex-gateway/pkg/graphql"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input domain.CreatePostInput) (string, error) {
	input.AuthorID = ctx.Value(domain.UserCtx).(string)

	// Create post.
	id, err := r.service.Post.CreatePost(ctx, input)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	err := r.service.Post.DeletePost(ctx, id, ctx.Value(domain.UserCtx).(string))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input domain.UpdatePostInput) (bool, error) {
	input.AuthorID = ctx.Value(domain.UserCtx).(string)

	// Update post.
	if err := r.service.Post.UpdatePost(ctx, input); err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*domain.Post, error) {
	// Getting a post.
	post, err := r.service.Post.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}

	// Getting author selections fields.
	fields := graphql.GetSelectionsFields(ctx, "author")

	// Check author selections fields.
	if len(fields) == 1 && fields[0] != "id" {
		// Getting post author.
		user, err := r.service.User.GetUserByID(ctx, post.Author.ID)
		if err != nil {
			return nil, err
		}

		post.Author = user
	}

	return post, nil
}
