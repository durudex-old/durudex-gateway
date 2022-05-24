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

	"github.com/durudex/durudex-gateway/internal/domain"
	v1 "github.com/durudex/durudex-gateway/pkg/pb/durudex/v1"
)

// User auth interface.
type Auth interface {
	SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error)
	SignOut(ctx context.Context, input domain.RefreshTokenInput) error
	RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error)
}

// User auth service structure.
type AuthService struct{ client v1.AuthServiceClient }

// Creating a new auth service.
func NewAuthService(client v1.AuthServiceClient) *AuthService {
	return &AuthService{client: client}
}

// User Sign In.
func (s *AuthService) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	response, err := s.client.UserSignIn(ctx, &v1.UserSignInRequest{
		Username: input.Username,
		Password: input.Password,
		Ip:       input.IP,
	})
	if err != nil {
		return nil, err
	}

	return &domain.Tokens{Access: response.Access, Refresh: response.Refresh}, nil
}

// User Sign Out.
func (s *AuthService) SignOut(ctx context.Context, input domain.RefreshTokenInput) error {
	_, err := s.client.UserSignOut(ctx, &v1.UserSignOutRequest{
		Refresh: input.Token,
		Ip:      input.IP,
	})

	return err
}

// Refresh user auth tokens by refresh token.
func (s *AuthService) RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error) {
	response, err := s.client.RefreshUserToken(ctx, &v1.RefreshUserTokenRequest{
		Refresh: input.Token,
		Ip:      input.IP,
	})
	if err != nil {
		return "", err
	}

	return response.Access, nil
}
