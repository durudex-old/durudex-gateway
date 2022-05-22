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

	"github.com/gofrs/uuid"
)

// User auth interface.
type Auth interface {
	SignUp(ctx context.Context, input domain.SignUpInput) (uuid.UUID, error)
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

// User Sign Up.
func (s *AuthService) SignUp(ctx context.Context, input domain.SignUpInput) (uuid.UUID, error) {
	response, err := s.client.UserSignUp(ctx, &v1.UserSignUpRequest{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.FromBytesOrNil(response.Id), nil
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
