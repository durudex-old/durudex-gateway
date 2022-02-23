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
	GetCodeByEmail(ctx context.Context, input domain.GetCodeByEmailInput) (*domain.Status, error)
}

// Code service structure.
type CodeService struct{ grpcHandler pb.CodeServiceClient }

// Creating a new code service.
func NewCodeService(grpcHandler pb.CodeServiceClient) *CodeService {
	return &CodeService{grpcHandler: grpcHandler}
}

// Getting code by email address.
func (s *CodeService) GetCodeByEmail(ctx context.Context, input domain.GetCodeByEmailInput) (*domain.Status, error) {
	status, err := s.grpcHandler.CreateByEmail(ctx, &types.Email{Email: input.Email})
	if err != nil {
		return &domain.Status{Status: status.Status}, err
	}

	return &domain.Status{Status: status.Status}, nil
}
