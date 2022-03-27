package graphql

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GraphQL directive for user authorization.
func (h *Handler) isAuth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if ctx.Value(domain.UserCtx) == nil {
		return nil, &gqlerror.Error{Message: "Authorization token failed"}
	}

	return next(ctx)
}
