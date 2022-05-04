package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) SignUp(ctx context.Context, input domain.SignUpInput) (string, error) {
	// Checking user input code.
	verify, err := r.service.User.VerifyEmailCode(ctx, input.Email, input.Code)
	if err != nil {
		return "", err
	} else if !verify {
		return "", &gqlerror.Error{
			Message:    "Invalid Code",
			Extensions: map[string]interface{}{"code": domain.CodeInvalidArgument},
		}
	}

	return r.service.Auth.SignUp(ctx, input)
}

func (r *mutationResolver) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	input.IP = ctx.Value(domain.IPCtx).(string)

	return r.service.Auth.SignIn(ctx, input)
}

func (r *mutationResolver) SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error) {
	input.IP = ctx.Value(domain.IPCtx).(string)

	return r.service.Auth.SignOut(ctx, input)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error) {
	input.IP = ctx.Value(domain.IPCtx).(string)

	return r.service.Auth.RefreshToken(ctx, input)
}
