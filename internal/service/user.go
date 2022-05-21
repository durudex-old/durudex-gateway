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

	"github.com/durudex/durudex-gateway/internal/domain"
)

// User interface.
type User interface {
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error)
	CreateVerifyEmailCode(ctx context.Context, email string) (bool, error)
	VerifyEmailCode(ctx context.Context, email string, code uint64) (bool, error)
}

// User service structure.
type UserService struct{}

// Creating a new user service.
func NewUserService() *UserService {
	return &UserService{}
}

// Get user by id.
func (s *UserService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}

// Forgot user password.
func (s *UserService) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) (bool, error) {
	return true, nil
}

// Creating a new verify user email code.
func (s *UserService) CreateVerifyEmailCode(ctx context.Context, email string) (bool, error) {
	return true, nil
}

// Verifying user email code.
func (s *UserService) VerifyEmailCode(ctx context.Context, email string, code uint64) (bool, error) {
	return true, nil
}
