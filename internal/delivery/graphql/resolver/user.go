package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/durudex/durudex-gateway/internal/domain"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateVerifyEmailCode(ctx context.Context, email string) (bool, error) {
	return r.service.User.CreateVerifyEmailCode(ctx, email)
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error) {
	// Checking user input code.
	status, err := r.service.User.VerifyEmailCode(ctx, input.Email, input.Code)
	if err != nil || !status {
		return false, &gqlerror.Error{Message: "Error sending email"}
	}

	return r.service.User.ForgotPassword(ctx, input)
}

func (r *mutationResolver) UpdateAvatar(ctx context.Context, file graphql.Upload) (string, error) {
	return "", nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*domain.User, error) {
	return r.service.User.GetUserByID(ctx, id)
}
