package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) SignUp(ctx context.Context, input domain.SignUpInput) (string, error) {
	// TODO: Add a check to the user service
	verify, err := r.service.User.VerifyEmailCode(ctx, input.Email, input.Code)
	if err != nil {
		return "", err
	} else if !verify {
		// Return error if email verification code is invalid.
		return "", &gqlerror.Error{
			Message:    "Invalid Code",
			Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
		}
	}

	// Sign Up.
	id, err := r.service.Auth.SignUp(ctx, input)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

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
	return r.service.Auth.RefreshToken(ctx, input)
}
