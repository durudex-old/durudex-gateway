/*
 * Copyright © 2022 Durudex

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
	"github.com/durudex/durudex-gateway/internal/config"
	"github.com/gofiber/fiber/v2"
)

// HTTP handler structure.
type Handler struct {
	cfg config.JWTConfig
}

// Creating a new HTTP handler.
func NewHandler(cfg config.JWTConfig) *Handler {
	return &Handler{cfg: cfg}
}

// Initialize http routes.
func (h *Handler) InitRoutes(router fiber.Router) {
	// Set auth middleware.
	router.Use(h.authMiddleware)

	// Ping pong route.
	router.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendString("pong")
	})
}