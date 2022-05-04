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

	"github.com/gofrs/uuid"
)

// User auth interface.
type Auth interface {
	SignUp(ctx context.Context, input domain.SignUpInput) (string, error)
	SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error)
	SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error)
	RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error)
}

// User auth service structure.
type AuthService struct{ grpcHandler pb.AuthServiceClient }

// Creating a new auth service.
func NewAuthService(grpcHandler pb.AuthServiceClient) *AuthService {
	return &AuthService{grpcHandler: grpcHandler}
}

// Sign Up user.
func (s *AuthService) SignUp(ctx context.Context, input domain.SignUpInput) (string, error) {
	// Sign Up user.
	response, err := s.grpcHandler.SignUp(ctx, &pb.SignUpRequest{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return "", err
	}

	// Get user uuid from bytes.
	id, err := uuid.FromBytes(response.Id)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

// Sign In user.
func (s *AuthService) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	tokens, err := s.grpcHandler.SignIn(ctx, &pb.SignInRequest{
		Username: input.Username,
		Password: input.Password,
		Ip:       input.IP,
	})
	if err != nil {
		return nil, err
	}

	return &domain.Tokens{Access: tokens.Access, Refresh: tokens.Refresh}, nil
}

// Sign Out user.
func (s *AuthService) SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error) {
	_, err := s.grpcHandler.SignOut(ctx, &pb.SignOutRequest{
		RefreshToken: input.Token,
		Ip:           input.IP,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

// Refresh user auth tokens by refresh token.
func (s *AuthService) RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error) {
	token, err := s.grpcHandler.RefreshToken(ctx, &pb.RefreshTokenRequest{
		RefreshToken: input.Token,
		Ip:           input.IP,
	})
	if err != nil {
		return "", err
	}

	return token.Access, nil
}
