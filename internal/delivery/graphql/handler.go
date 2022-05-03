/*
 * Copyright © 2021-2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package graphql

import (
	"context"
	"net/http"

	"github.com/durudex/durudex-gateway/internal/delivery/graphql/generated"
	"github.com/durudex/durudex-gateway/internal/delivery/graphql/resolver"
	"github.com/durudex/durudex-gateway/internal/domain"
	"github.com/durudex/durudex-gateway/internal/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GraphQL handler structure.
type Handler struct{ service *service.Service }

// Creating a new graphql handler.
func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// GraphQL handler.
func (h *Handler) GraphqlHandler() http.HandlerFunc {
	// GraphQL config.
	config := generated.Config{
		Resolvers:  resolver.NewResolver(h.service),
		Directives: generated.DirectiveRoot{IsAuth: h.isAuth},
	}

	// Creating a new graphql handler.
	handler := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	// Set graphql error handler.
	handler.SetErrorPresenter(h.errorHandler)
	// Set graphql panic handler.
	handler.SetRecoverFunc(h.recoverHandler)

	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

// GraphQL playground handler.
func (h *Handler) PlaygroundHandler() http.HandlerFunc {
	return playground.Handler("GraphQL", "/graph/query")
}

// GraphQL error handler.
func (h *Handler) errorHandler(ctx context.Context, err error) *gqlerror.Error {
	return &gqlerror.Error{}
}

// GraphQL recover handler.
func (h *Handler) recoverHandler(ctx context.Context, err interface{}) error {
	return &gqlerror.Error{
		Message:    "Internal server error",
		Extensions: map[string]interface{}{"code": domain.CodeInternalServerError},
	}
}
