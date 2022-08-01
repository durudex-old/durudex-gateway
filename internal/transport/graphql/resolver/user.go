package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/durudex/durudex-gateway/internal/domain"
	"github.com/durudex/durudex-gateway/internal/transport/graphql/generated"
	"github.com/segmentio/ksuid"
)

// CreateVerifyEmailCode is the resolver for the createVerifyEmailCode field.
func (r *mutationResolver) CreateVerifyEmailCode(ctx context.Context, email string) (bool, error) {
	if err := r.service.User.CreateVerifyEmailCode(ctx, email); err != nil {
		return false, err
	}

	return true, nil
}

// ForgotPassword is the resolver for the forgotPassword field.
func (r *mutationResolver) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error) {
	if err := r.service.User.ForgotPassword(ctx, input); err != nil {
		return false, err
	}

	return true, nil
}

// UpdateAvatar is the resolver for the updateAvatar field.
func (r *mutationResolver) UpdateAvatar(ctx context.Context, file graphql.Upload) (string, error) {
	return "", nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*domain.User, error) {
	// Parsing id from string.
	id, err := ksuid.Parse(ctx.Value(domain.UserCtx).(string))
	if err != nil {
		return nil, err
	}

	// Getting user.
	user, err := r.service.User.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id ksuid.KSUID) (*domain.User, error) {
	user, err := r.service.User.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *domain.User, first *int, last *int) (*domain.PostConnection, error) {
	return &domain.PostConnection{}, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
