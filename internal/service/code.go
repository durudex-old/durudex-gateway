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

	"github.com/durudex/durudex-gateway/internal/delivery/grpc/pb"
	"github.com/durudex/durudex-gateway/internal/delivery/grpc/pb/types"
	"github.com/durudex/durudex-gateway/internal/domain"
)

// Code service interface.
type Code interface {
	GetByEmail(ctx context.Context, input domain.GetCodeByEmailInput) (bool, error)
	CheckByEmail(ctx context.Context, email string, input uint64) (bool, error)
}

// Code service structure.
type CodeService struct{ grpcHandler pb.CodeServiceClient }

// Creating a new code service.
func NewCodeService(grpcHandler pb.CodeServiceClient) *CodeService {
	return &CodeService{grpcHandler: grpcHandler}
}

// Getting code by email address.
func (s *CodeService) GetByEmail(ctx context.Context, input domain.GetCodeByEmailInput) (bool, error) {
	status, err := s.grpcHandler.CreateByEmail(ctx, &types.Email{Email: input.Email})
	if err != nil {
		return status.Status, err
	}

	return status.Status, nil
}

func (s *CodeService) CheckByEmail(ctx context.Context, email string, input uint64) (bool, error) {
	code, err := s.grpcHandler.GetByEmail(ctx, &types.Email{Email: email})
	if err != nil {
		return false, err
	}

	if code.Code != input {
		return false, nil // todo
	}

	return true, nil
}
