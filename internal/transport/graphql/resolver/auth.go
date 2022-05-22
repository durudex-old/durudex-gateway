package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/durudex/durudex-gateway/internal/domain"
)

func (r *mutationResolver) SignUp(ctx context.Context, input domain.SignUpInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
