package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
)

func (r *mutationResolver) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	input.IP = ctx.Value(domain.IPCtx).(string)

	// Sign In.
	return r.service.Auth.SignIn(ctx, input)
}

func (r *mutationResolver) SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error) {
	input.IP = ctx.Value(domain.IPCtx).(string)

	// Sign Out.
	if err := r.service.Auth.SignOut(ctx, input); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error) {
	input.IP = ctx.Value(domain.IPCtx).(string)

	// Refresh token.
	return r.service.Auth.RefreshToken(ctx, input)
}
