/*
	Copyright © 2021-2022 Durudex

	This file is part of Durudex: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	Durudex is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with Durudex. If not, see <https://www.gnu.org/licenses/>.
*/

package graphql

import (
	"context"

	"github.com/Durudex/durudex-gateway/internal/delivery/graphql/model"
	"github.com/Durudex/durudex-gateway/internal/delivery/grpc/pb"
	"github.com/Durudex/durudex-gateway/internal/delivery/grpc/pb/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Sign Up resolver.
func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpInput) (*model.SignUp, error) {
	// Sign Up user for auth service.
	user := pb.SignUpRequest{
		Username: input.Username,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Birthday: timestamppb.New(input.Birthday),
		Sex:      input.Sex,
	}
	id, err := r.service.Auth.SignUp(ctx, &user)

	if err != nil {
		return &model.SignUp{}, err
	}

	return &model.SignUp{ID: id}, nil
}

// Sign In resolver.
func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.SignIn, error) {
	// Sign In user for auth service.
	user := pb.SignInRequest{
		Username: input.Username,
		Password: input.Password,
		Ip:       ctx.Value(userIP).(string),
	}
	tokens, err := r.service.Auth.SignIn(ctx, &user)

	if err != nil {
		return &model.SignIn{}, err
	}

	return &model.SignIn{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

// Refresh user auth tokens.
func (r *mutationResolver) RefreshTokens(ctx context.Context, input model.RefreshTokensInput) (*model.RefreshTokens, error) {
	// Refresh auth token for auth service.
	refreshToken := pb.RefreshTokensRequest{
		RefreshToken: input.RefreshToken,
		Ip:           ctx.Value(userIP).(string),
	}
	tokens, err := r.service.Auth.RefreshTokens(ctx, &refreshToken)

	if err != nil {
		return &model.RefreshTokens{}, err
	}

	return &model.RefreshTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (r *mutationResolver) Verify(ctx context.Context, input model.VerifyInput) (*model.Status, error) {
	// Get user id by ctx.
	userID, err := GetUserID(ctx)
	if err != nil {
		return &model.Status{Status: false}, err
	}

	// Verify user.
	request := pb.VerifyRequest{
		Id:   userID,
		Code: input.Code,
	}

	status, err := r.service.Auth.Verify(ctx, &request)
	if err != nil {
		return &model.Status{Status: status}, err
	}

	return &model.Status{Status: status}, nil
}

// Get verify code in user email.
func (r *queryResolver) GetCode(ctx context.Context) (*model.Status, error) {
	// Get user id by ctx.
	userID, err := GetUserID(ctx)
	if err != nil {
		return &model.Status{Status: false}, err
	}

	status, err := r.service.Auth.GetCode(ctx, &types.ID{Id: userID})
	if err != nil {
		return &model.Status{Status: status}, err
	}

	return &model.Status{Status: status}, nil
}
