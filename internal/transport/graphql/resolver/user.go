package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/durudex/durudex-gateway/internal/domain"
)

func (r *mutationResolver) SignUp(ctx context.Context, input domain.SignUpInput) (string, error) {
	id, err := r.service.User.SignUp(ctx, input)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *mutationResolver) CreateVerifyEmailCode(ctx context.Context, email string) (bool, error) {
	if err := r.service.User.CreateVerifyEmailCode(ctx, email); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error) {
	if err := r.service.User.ForgotPassword(ctx, input); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateAvatar(ctx context.Context, file graphql.Upload) (string, error) {
	return "", nil
}

func (r *queryResolver) Me(ctx context.Context) (*domain.User, error) {
	user, err := r.service.User.GetUserByID(ctx, ctx.Value(domain.UserCtx).(string))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*domain.User, error) {
	user, err := r.service.User.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
