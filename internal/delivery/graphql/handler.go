/*
	Copyright © 2021 Durudex

	This file is part of Durudex: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	Durudex is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with Durudex. If not, see <https://www.gnu.org/licenses/>.
*/

package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Durudex/durudex-auth-service/pkg/auth"
	"github.com/Durudex/durudex-gateway/internal/delivery/graphql/generated"
	"github.com/Durudex/durudex-gateway/internal/service"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *service.Service
	auth    *auth.Manager
}

// Creating a new graphql handler.
func NewGraphQLHandler(service *service.Service, auth *auth.Manager) *Handler {
	return &Handler{
		service: service,
		auth:    auth,
	}
}

// Defining the graphql handler.
func (h *Handler) graphqlHandler() http.HandlerFunc {
	// NewExecutableSchema and Config are in the generate.go file.
	// Resolver is in the resolver.go file.
	config := generated.Config{
		Resolvers: NewResolver(h.service),
	}
	config.Directives.UserAuth = h.userAuth

	handler := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

// Defining the Playground handler.
func (h *Handler) playgroundHandler() http.HandlerFunc {
	handler := playground.Handler("GraphQL", "/query")

	return func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}
}

// Initialize graphql routes.
func (h *Handler) InitRoutes(router fiber.Router) {
	router.Use(h.authMiddleware)
	router.Post("/query", adaptor.HTTPHandlerFunc(h.graphqlHandler()))
	router.Get("/", adaptor.HTTPHandlerFunc(h.playgroundHandler()))
}
