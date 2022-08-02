/*
 * Copyright © 2022 Durudex
 *
 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package http

import (
	"github.com/durudex/durudex-gateway/internal/config"
	"github.com/durudex/durudex-gateway/internal/transport/graphql"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// HTTP handler structure.
type Handler struct {
	config     *config.HTTPConfig
	graphql    *graphql.Handler
	signingKey string
}

// Creating a new HTTP handler.
func NewHandler(config *config.HTTPConfig, graphql *graphql.Handler, signingKey string) *Handler {
	return &Handler{config: config, graphql: graphql, signingKey: signingKey}
}

// Initialize http middleware.
func (h *Handler) InitMiddleware(router fiber.Router) {
	if h.config.Cors.Enable {
		// CORS configuration.
		corsConfig := cors.Config{
			AllowOrigins: h.config.Cors.AllowOrigins,
			AllowMethods: h.config.Cors.AllowMethods,
			AllowHeaders: h.config.Cors.AllowHeaders,
		}

		// Initialize middleware.
		router.Use(cors.New(corsConfig))
	}
}

// Initialize http routes.
func (h *Handler) InitRoutes(router fiber.Router) {
	// Ping pong route.
	router.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})

	graph := router.Group("/graph", h.authMiddleware)
	{
		graph.Get("/", adaptor.HTTPHandlerFunc(h.graphql.PlaygroundHandler()))
		graph.Post("/query", adaptor.HTTPHandlerFunc(h.graphql.GraphqlHandler()))
	}
}
