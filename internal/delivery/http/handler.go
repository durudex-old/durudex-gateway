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

package http

import (
	"github.com/Durudex/durudex-gateway/internal/delivery/graphql"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	graphqlHandler *graphql.Handler
}

// Creating a new http handler
func NewHTTPHandler(graphqlHandler *graphql.Handler) *Handler {
	return &Handler{
		graphqlHandler: graphqlHandler,
	}
}

// Initialize http routes.
func (h *Handler) InitRoutes(router fiber.Router) {
	// Ping pong route
	router.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	// GrapgQL routes.
	h.graphqlHandler.InitRoutes(router)
}
