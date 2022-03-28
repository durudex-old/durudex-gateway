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

package service

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/delivery/grpc/pb"
	"github.com/durudex/durudex-gateway/internal/domain"
)

// User auth interface.
type Auth interface {
	SignUp(ctx context.Context, input domain.SignUpInput) (uint64, error)
	SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error)
	RefreshTokens(ctx context.Context, input domain.RefreshTokensInput) (*domain.Tokens, error)
	Logout(ctx context.Context, input domain.RefreshTokensInput) (bool, error)
}

// User auth service structure.
type AuthService struct{ grpcHandler pb.AuthUserServiceClient }

// Creating a new auth service.
func NewAuthService(grpcHandler pb.AuthUserServiceClient) *AuthService {
	return &AuthService{grpcHandler: grpcHandler}
}

// Sign Up user.
func (s *AuthService) SignUp(ctx context.Context, input domain.SignUpInput) (uint64, error) {
	id, err := s.grpcHandler.SignUp(ctx, &pb.SignUpRequest{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return 0, err
	}

	return id.Id, err
}

// Sign In user.
func (s *AuthService) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	tokens, err := s.grpcHandler.SignIn(ctx, &pb.SignInRequest{
		Username: input.Username,
		Password: input.Password,
		Ip:       input.IP,
	})
	if err != nil {
		return &domain.Tokens{}, err
	}

	return &domain.Tokens{Access: tokens.Access, Refresh: tokens.Refresh}, nil
}

// Refresh user auth tokens by refresh token.
func (s *AuthService) RefreshTokens(ctx context.Context, input domain.RefreshTokensInput) (*domain.Tokens, error) {
	tokens, err := s.grpcHandler.RefreshTokens(ctx, &pb.RefreshTokenRequest{
		RefreshToken: input.Token,
		Ip:           input.IP,
	})
	if err != nil {
		return &domain.Tokens{}, err
	}

	return &domain.Tokens{Access: tokens.Access, Refresh: tokens.Refresh}, nil
}

// Logout user session by refresh token.
func (s *AuthService) Logout(ctx context.Context, input domain.RefreshTokensInput) (bool, error) {
	status, err := s.grpcHandler.Logout(ctx, &pb.RefreshTokenRequest{
		RefreshToken: input.Token,
		Ip:           input.IP,
	})
	if err != nil {
		return status.Status, err
	}

	return status.Status, nil
}
