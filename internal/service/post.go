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
	"github.com/durudex/durudex-gateway/internal/domain"

	"github.com/gofrs/uuid"
)

// Post service interface.
type Post interface {
	CreatePost(ctx context.Context, input domain.CreatePostInput) (string, error)
	DeletePost(ctx context.Context, id string) (bool, error)
	GetPost(ctx context.Context, id string) (*domain.Post, error)
}

// Post service structure.
type PostService struct{ grpcHandler pb.PostServiceClient }

// Creating a new post service.
func NewPostService(grpcHandler pb.PostServiceClient) *PostService {
	return &PostService{grpcHandler: grpcHandler}
}

// Creating a new post.
func (s *PostService) CreatePost(ctx context.Context, input domain.CreatePostInput) (string, error) {
	// Create post.
	id, err := s.grpcHandler.CreatePost(ctx, &pb.CreatePostRequest{Text: input.Text})
	if err != nil {
		return "", err
	}

	// Get user uuid from bytes.
	userID, err := uuid.FromBytes(id.Id)
	if err != nil {
		return "", err
	}

	return userID.String(), err
}

// Deleting a post.
func (s *PostService) DeletePost(ctx context.Context, id string) (bool, error) {
	// Get user uuid from string.
	userID, err := uuid.FromString(id)
	if err != nil {
		return false, err
	}

	// Delete post.
	_, err = s.grpcHandler.DeletePost(ctx, &pb.DeletePostRequest{Id: userID.Bytes()})
	if err != nil {
		return false, err
	}

	return true, nil
}

// Getting a post.
func (s *PostService) GetPost(ctx context.Context, id string) (*domain.Post, error) {
	// Get user uuid from string.
	userID, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	// Get post by id.
	post, err := s.grpcHandler.GetPostByID(ctx, &pb.GetPostByIDRequest{Id: userID.Bytes()})
	if err != nil {
		return nil, err
	}

	updatedAt := post.UpdatedAt.AsTime()

	return &domain.Post{
		ID:        id,
		AuthorID:  uuid.FromBytesOrNil(post.AuthorId).String(),
		Text:      post.Text,
		CreatedAt: post.CreatedAt.AsTime(),
		UpdatedAt: &updatedAt,
	}, nil
}
