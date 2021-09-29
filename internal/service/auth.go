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

package service

import (
	"context"

	"github.com/Durudex/durudex-gateway/internal/delivery/grpc"
	pb "github.com/Durudex/durudex-gateway/internal/delivery/grpc/protobuf"
)

type AuthService struct {
	grpcHandler *grpc.Handler
}

// Creating a new auth service.
func NewAuthService(grpcHandler *grpc.Handler) *AuthService {
	return &AuthService{
		grpcHandler: grpcHandler,
	}
}

// Sign Up user.
func (s *AuthService) SignUp(ctx context.Context, input *pb.UserSignUpRequest) (uint64, error) {
	// Get for auth service.
	id, err := s.grpcHandler.Auth.SignUp(ctx, input)
	if err != nil {
		return 0, err
	}

	return id.Id, err
}
