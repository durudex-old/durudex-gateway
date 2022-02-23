package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
)

func (r *mutationResolver) GetCodeByEmail(ctx context.Context, input domain.GetCodeByEmailInput) (*domain.Status, error) {
	return r.service.Code.GetCodeByEmail(ctx, input)
}
