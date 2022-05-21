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
)

// User auth interface.
type Auth interface {
	SignUp(ctx context.Context, input domain.SignUpInput) (string, error)
	SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error)
	SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error)
	RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error)
}

// User auth service structure.
type AuthService struct{}

// Creating a new auth service.
func NewAuthService() *AuthService {
	return &AuthService{}
}

// Sign Up user.
func (s *AuthService) SignUp(ctx context.Context, input domain.SignUpInput) (string, error) {
	return "", nil
}

// Sign In user.
func (s *AuthService) SignIn(ctx context.Context, input domain.SignInInput) (*domain.Tokens, error) {
	return nil, nil
}

// Sign Out user.
func (s *AuthService) SignOut(ctx context.Context, input domain.RefreshTokenInput) (bool, error) {
	return true, nil
}

// Refresh user auth tokens by refresh token.
func (s *AuthService) RefreshToken(ctx context.Context, input domain.RefreshTokenInput) (string, error) {
	return "", nil
}
