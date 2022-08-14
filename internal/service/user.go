/*
 * Copyright © 2022 Durudex
 *
 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package service

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"
	v1 "github.com/durudex/durudex-gateway/pkg/pb/durudex/v1"

	"github.com/segmentio/ksuid"
)

// User interface.
type User interface {
	GetById(ctx context.Context, id ksuid.KSUID) (*domain.User, error)
	ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) error
	CreateVerifyEmailCode(ctx context.Context, email string) error
	VerifyEmailCode(ctx context.Context, email string, code uint64) (bool, error)
}

// User service structure.
type UserService struct {
	user v1.UserServiceClient
	code v1.UserCodeServiceClient
}

// Creating a new user service.
func NewUserService(user v1.UserServiceClient, code v1.UserCodeServiceClient) *UserService {
	return &UserService{user: user, code: code}
}

// Getting user by id.
func (s *UserService) GetById(ctx context.Context, id ksuid.KSUID) (*domain.User, error) {
	response, err := s.user.GetUserById(ctx, &v1.GetUserByIdRequest{Id: id.Bytes()})

	return &domain.User{
		Id:        id,
		Username:  response.Username,
		LastVisit: response.LastVisit.AsTime(),
		Verified:  response.Verified,
		AvatarUrl: response.AvatarUrl,
	}, err
}

// Forgot user password.
func (s *UserService) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) error {
	_, err := s.user.ForgotUserPassword(ctx, &v1.ForgotUserPasswordRequest{
		Email:    input.Email,
		Password: input.Password,
		Code:     input.Code,
	})

	return err
}

// Creating a new verify user email code.
func (s *UserService) CreateVerifyEmailCode(ctx context.Context, email string) error {
	_, err := s.code.CreateVerifyUserEmailCode(ctx, &v1.CreateVerifyUserEmailCodeRequest{
		Email: email,
	})

	return err
}

// Verifying user email code.
func (s *UserService) VerifyEmailCode(ctx context.Context, email string, code uint64) (bool, error) {
	response, err := s.code.VerifyUserEmailCode(ctx, &v1.VerifyUserEmailCodeRequest{
		Email: email,
		Code:  code,
	})
	if err != nil {
		return false, err
	}

	return response.Status, nil
}
