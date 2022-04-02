package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
)

func (r *mutationResolver) SignUp(ctx context.Context, input domain.SignUpInput) (uint64, error) {
	// Checking user input code.
	status, err := r.service.Code.CheckByEmail(ctx, input.Email, input.Code)
	if err != nil || !status {
		return 0, err
	}

	return r.service.Auth.SignUp(ctx, input)
}

func (r *mutationResolver) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	input.IP = ctx.Value(domain.IPCtx).(string)

	return r.service.Auth.SignIn(ctx, input)
}

func (r *mutationResolver) RefreshTokens(ctx context.Context, input domain.RefreshTokenInput) (*domain.Tokens, error) {
	input.IP = ctx.Value(domain.IPCtx).(string)

	return r.service.Auth.RefreshTokens(ctx, input)
}

func (r *mutationResolver) Logout(ctx context.Context, input domain.RefreshTokenInput) (bool, error) {
	input.IP = ctx.Value(domain.IPCtx).(string)

	return r.service.Auth.Logout(ctx, input)
}
