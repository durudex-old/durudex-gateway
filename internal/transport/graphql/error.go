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

package graphql

import (
	"context"
	"errors"

	"github.com/durudex/durudex-gateway/internal/domain"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Getting gRPC status code from gRPC error.
func fromGRPCError(err error) (*status.Status, bool) {
	// Check if error is gRPC error.
	if se, ok := err.(interface {
		GRPCStatus() *status.Status
	}); ok {
		// Return gRPC status code.
		return se.GRPCStatus(), true
	} else {
		return nil, false
	}
}

// GraphQL error handler.
func (h *Handler) errorHandler(ctx context.Context, err error) *gqlerror.Error {
	var gqlErr *gqlerror.Error

	// Check if error is a gqlerror.Error.
	if errors.As(err, &gqlErr) {
		// Get gRPC status code from error.
		if e, ok := fromGRPCError(gqlErr.Unwrap()); ok {
			// GRPC error handler.
			return h.grpcErrorHandler(e.Code())
		} else {
			return gqlErr
		}
	} else {
		// Default graphql error.
		return &gqlerror.Error{
			Message:    "Server error",
			Extensions: map[string]interface{}{"code": domain.CodeServerError},
		}
	}
}

// gRPC error handler.
func (h *Handler) grpcErrorHandler(code codes.Code) *gqlerror.Error {
	switch code {
	case codes.Internal:
		// Set internal server error.
		return &gqlerror.Error{
			Message:    "Internal server error",
			Extensions: map[string]interface{}{"code": domain.CodeInternalServerError},
		}
	default:
		// Default grpc error.
		return &gqlerror.Error{
			Message:    "Server error",
			Extensions: map[string]interface{}{"code": domain.CodeServerError},
		}
	}
}

// GraphQL recover handler.
func (h *Handler) recoverHandler(ctx context.Context, err interface{}) error {
	return &gqlerror.Error{
		Message:    "Internal server error",
		Extensions: map[string]interface{}{"code": domain.CodeInternalServerError},
	}
}
