package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateVerifyEmailCode(ctx context.Context, email string) (bool, error) {
	if err := r.service.User.CreateVerifyEmailCode(ctx, email); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error) {
	// TODO: Add a check to the user service
	verify, err := r.service.User.VerifyEmailCode(ctx, input.Email, input.Code)
	if err != nil {
		return false, err
	} else if !verify {
		// Return error if email verification code is invalid.
		return false, &gqlerror.Error{
			Message:    "Invalid Code",
			Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
		}
	}

	// Forgot password.
	if err := r.service.User.ForgotPassword(ctx, input); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) UpdateAvatar(ctx context.Context, file graphql.Upload) (string, error) {
	return "", nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*domain.User, error) {
	user, err := r.service.User.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
