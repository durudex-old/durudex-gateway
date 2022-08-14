package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input domain.SignUpInput) (*domain.Tokens, error) {
	input.Ip = ctx.Value(domain.IpCtx).(string)

	// Sign Up.
	return r.service.Auth.SignUp(ctx, input)
}

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	input.Ip = ctx.Value(domain.IpCtx).(string)

	// Sign In.
	return r.service.Auth.SignIn(ctx, input)
}

// SignOut is the resolver for the signOut field.
func (r *mutationResolver) SignOut(ctx context.Context, token string) (bool, error) {
	// Sign Out.
	if err := r.service.Auth.SignOut(ctx, domain.RefreshTokenInput{
		Token: token,
		Ip:    ctx.Value(domain.IpCtx).(string),
	}); err != nil {
		return false, err
	}

	return true, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, token string) (string, error) {
	// Refresh token.
	return r.service.Auth.RefreshToken(ctx, domain.RefreshTokenInput{
		Token: token,
		Ip:    ctx.Value(domain.IpCtx).(string),
	})
}
