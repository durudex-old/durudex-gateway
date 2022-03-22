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

package http

import (
	"github.com/durudex/durudex-gateway/internal/delivery/graphql"
	"github.com/durudex/durudex-gateway/internal/service"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// HTTP handler structure.
type Handler struct{ service *service.Service }

// Creating a new http handler.
func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// Initialize http routes.
func (h *Handler) InitRoutes(router fiber.Router) {
	router.Use(h.authMiddleware)

	router.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	graphql := graphql.NewHandler(h.service)

	graph := router.Group("/graph")
	{
		graph.Get("/", adaptor.HTTPHandlerFunc(graphql.PlaygroundHandler()))
		graph.Post("/query", adaptor.HTTPHandlerFunc(graphql.GraphqlHandler()))
	}
}
