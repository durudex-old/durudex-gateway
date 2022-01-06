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
	"github.com/Durudex/durudex-gateway/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	grpcHandler *grpc.Handler
}

// Creating a new auth service.
func NewAuthService(grpcHandler *grpc.Handler) *AuthService {
	return &AuthService{grpcHandler: grpcHandler}
}

// Sign Up user.
func (s *AuthService) SignUp(ctx context.Context, input *domain.SignUpInput) (uint64, error) {
	// Get for auth service.
	user := pb.SignUpRequest{
		Username: input.Username,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Birthday: timestamppb.New(input.Birthday),
		Sex:      input.Sex,
	}
	id, err := s.grpcHandler.Auth.SignUp(ctx, &user)
	if err != nil {
		return 0, err
	}

	return id.Id, err
}

// Sign In user.
func (s *AuthService) SignIn(ctx context.Context, input *domain.SignInInput) (domain.Tokens, error) {
	// Get for auth service.
	user := pb.SignInRequest{
		Username: input.Username,
		Password: input.Password,
		Ip:       input.Ip,
	}
	tokens, err := s.grpcHandler.Auth.SignIn(ctx, &user)
	if err != nil {
		return domain.Tokens{}, err
	}

	return domain.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

// Refresh user auth tokens.
func (s *AuthService) RefreshTokens(ctx context.Context, input *domain.RefreshTokensInput) (domain.Tokens, error) {
	// Get for auth service.
	refreshTokensInput := pb.RefreshTokensRequest{
		RefreshToken: input.RefreshToken,
		Ip:           input.Ip,
	}
	tokens, err := s.grpcHandler.Auth.RefreshTokens(ctx, &refreshTokensInput)
	if err != nil {
		return domain.Tokens{}, err
	}

	return domain.Tokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

// Verification user.
func (s *AuthService) Verify(ctx context.Context, input *domain.VerifyInput) (bool, error) {
	// Get for auth service.
	verifyInput := pb.VerifyRequest{Id: input.Id, Code: input.Code}
	verifyStatus, err := s.grpcHandler.Auth.Verify(ctx, &verifyInput)
	if err != nil {
		return verifyStatus.Status, err
	}

	return verifyStatus.Status, nil
}

// Get user verification code.
func (s *AuthService) GetCode(ctx context.Context, id uint64) (bool, error) {
	emailStatus, err := s.grpcHandler.Auth.GetCode(ctx, &types.ID{Id: id})
	if err != nil {
		return emailStatus.Status, err
	}

	return emailStatus.Status, nil
}
