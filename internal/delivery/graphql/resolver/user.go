package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
)

func (r *mutationResolver) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error) {
	// Checking user input code.
	status, err := r.service.Code.CheckByEmail(ctx, input.Email, input.Code)
	if err != nil || !status {
		return false, err
	}

	return r.service.User.ForgotPassword(ctx, input)
}
