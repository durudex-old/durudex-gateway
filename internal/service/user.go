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
	v1 "github.com/durudex/durudex-gateway/pkg/pb/durudex/v1"

	"github.com/gofrs/uuid"
)

// User interface.
type User interface {
	SignUp(ctx context.Context, input domain.SignUpInput) (uuid.UUID, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) error
	CreateVerifyEmailCode(ctx context.Context, email string) error
	VerifyEmailCode(ctx context.Context, email string, code uint64) (bool, error)
}

// User service structure.
type UserService struct{ client v1.UserServiceClient }

// Creating a new user service.
func NewUserService(client v1.UserServiceClient) *UserService {
	return &UserService{client: client}
}

// User Sign Up.
func (s *UserService) SignUp(ctx context.Context, input domain.SignUpInput) (uuid.UUID, error) {
	response, err := s.client.UserSignUp(ctx, &v1.UserSignUpRequest{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		Code:     input.Code,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.FromBytesOrNil(response.Id), nil
}

// Get user by id.
func (s *UserService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	// Getting user uuid from string.
	userID, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	// Getting user by id.
	response, err := s.client.GetUserById(ctx, &v1.GetUserByIdRequest{Id: userID.Bytes()})
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        id,
		Username:  response.Username,
		CreatedAt: response.CreatedAt.AsTime(),
		LastVisit: response.LastVisit.AsTime(),
		Verified:  response.Verified,
		AvatarURL: response.AvatarUrl,
	}, nil
}

// Forgot user password.
func (s *UserService) ForgotPassword(ctx context.Context, input domain.ForgotPasswordInput) error {
	_, err := s.client.ForgotUserPassword(ctx, &v1.ForgotUserPasswordRequest{
		Email:    input.Email,
		Password: input.Password,
		Code:     input.Code,
	})
	if err != nil {
		return err
	}

	return err
}

// Creating a new verify user email code.
func (s *UserService) CreateVerifyEmailCode(ctx context.Context, email string) error {
	_, err := s.client.CreateVerifyUserEmailCode(ctx, &v1.CreateVerifyUserEmailCodeRequest{
		Email: email,
	})

	return err
}

// Verifying user email code.
func (s *UserService) VerifyEmailCode(ctx context.Context, email string, code uint64) (bool, error) {
	response, err := s.client.VerifyUserEmailCode(ctx, &v1.VerifyUserEmailCodeRequest{
		Email: email,
		Code:  code,
	})
	if err != nil {
		return false, err
	}

	return response.Status, nil
}
