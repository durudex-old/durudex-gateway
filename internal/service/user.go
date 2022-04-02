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

package service

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/delivery/grpc/pb"
	"github.com/durudex/durudex-gateway/internal/domain"
)

// User interface.
type User interface {
	ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error)
}

// User service structure.
type UserService struct{ grpcHandler pb.UserServiceClient }

// Creating a new user service.
func NewUserService(grpcHandler pb.UserServiceClient) *UserService {
	return &UserService{grpcHandler: grpcHandler}
}

// Forgot user password.
func (s *UserService) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error) {
	status, err := s.grpcHandler.ForgotPassword(ctx, &pb.ForgotPasswordRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return status.Status, err
	}

	return status.Status, nil
}
