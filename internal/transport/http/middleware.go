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
	"strings"

	"github.com/durudex/durudex-gateway/internal/domain"
	"github.com/durudex/durudex-gateway/pkg/auth"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofiber/fiber/v2"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Authorization http header.
const authorizationHeader string = "Authorization"

// HTTP authorization middleware.
func (h *Handler) authMiddleware(ctx *fiber.Ctx) error {
	// Set ip to context value.
	ctx.Context().SetUserValue(domain.IpCtx, ctx.IP())

	// Getting authorization header.
	header := ctx.Get(authorizationHeader)
	if header == "" {
		return ctx.Next()
	}

	// Checking header parts.
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			graphql.Response{
				Errors: []*gqlerror.Error{
					{
						Message:    "Invalid authorization header",
						Extensions: map[string]interface{}{"code": domain.CodeUnauthorized},
					},
				},
			},
		)
	}

	// Parsing jwt access token.
	customClaim, err := auth.Parse(headerParts[1], h.signingKey)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			graphql.Response{
				Errors: []*gqlerror.Error{
					{
						Message:    "Authorization token is invalid",
						Extensions: map[string]interface{}{"code": domain.CodeUnauthorized},
					},
				},
			},
		)
	}

	ctx.Context().SetUserValue(domain.UserCtx, customClaim)

	return ctx.Next()
}
