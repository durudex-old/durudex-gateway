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

// Post service interface.
type Post interface {
	CreatePost(ctx context.Context, input domain.CreatePostInput) (ksuid.KSUID, error)
	DeletePost(ctx context.Context, id, authorId ksuid.KSUID) error
	UpdatePost(ctx context.Context, input domain.UpdatePostInput) error
	GetPost(ctx context.Context, id ksuid.KSUID) (*domain.Post, error)
}

// Post service structure.
type PostService struct{ client v1.PostServiceClient }

// Creating a new post service.
func NewPostService(client v1.PostServiceClient) *PostService {
	return &PostService{client: client}
}

// Creating a new post.
func (s *PostService) CreatePost(ctx context.Context, input domain.CreatePostInput) (ksuid.KSUID, error) {
	response, err := s.client.CreatePost(ctx, &v1.CreatePostRequest{
		AuthorId: input.AuthorId.Bytes(),
		Text:     input.Text,
	})
	if err != nil {
		return ksuid.Nil, err
	}

	return ksuid.FromBytes(response.Id)
}

// Deleting a post.
func (s *PostService) DeletePost(ctx context.Context, id, authorId ksuid.KSUID) error {
	if _, err := s.client.DeletePost(ctx, &v1.DeletePostRequest{
		Id:       id.Bytes(),
		AuthorId: authorId.Bytes(),
	}); err != nil {
		return err
	}

	return nil
}

// Updating a post.
func (s *PostService) UpdatePost(ctx context.Context, input domain.UpdatePostInput) error {
	if _, err := s.client.UpdatePost(ctx, &v1.UpdatePostRequest{
		Id:       input.Id.Bytes(),
		AuthorId: input.AuthorId.Bytes(),
		Text:     input.Text,
	}); err != nil {
		return err
	}

	return nil
}

// Getting a post.
func (s *PostService) GetPost(ctx context.Context, id ksuid.KSUID) (*domain.Post, error) {
	post, err := s.client.GetPostById(ctx, &v1.GetPostByIdRequest{Id: id.Bytes()})
	if err != nil {
		return nil, err
	}

	// Getting post author ksuid from bytes.
	authorId, err := ksuid.FromBytes(post.AuthorId)
	if err != nil {
		return nil, err
	}

	return &domain.Post{
		Id:        id,
		Author:    &domain.User{Id: authorId},
		Text:      post.Text,
		UpdatedAt: post.UpdatedAt.AsOptionalTime(),
	}, nil
}
