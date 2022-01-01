/*
	Copyright © 2021-2022 Durudex

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
	"github.com/Durudex/durudex-gateway/internal/delivery/grpc/pb"
	"github.com/Durudex/durudex-gateway/internal/delivery/grpc/pb/types"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Auth interface {
	SignUp(ctx context.Context, input *pb.SignUpRequest) (uint64, error)
	SignIn(ctx context.Context, input *pb.SignInRequest) (Tokens, error)
	Verify(context.Context, *pb.VerifyRequest) (bool, error)
	GetCode(context.Context, *types.ID) (bool, error)
	RefreshTokens(ctx context.Context, input *pb.RefreshTokensRequest) (Tokens, error)
}

type Service struct {
	Auth
}

// Creating a new service.
func NewService(grpcHandler *grpc.Handler) *Service {
	return &Service{Auth: NewAuthService(grpcHandler)}
}
