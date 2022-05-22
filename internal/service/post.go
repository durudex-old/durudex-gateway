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

// Post service interface.
type Post interface {
	CreatePost(ctx context.Context, input domain.CreatePostInput) (uuid.UUID, error)
	DeletePost(ctx context.Context, id string) error
	UpdatePost(ctx context.Context, input domain.UpdatePostInput) error
	GetPost(ctx context.Context, id string) (*domain.Post, error)
}

// Post service structure.
type PostService struct{ client v1.PostServiceClient }

// Creating a new post service.
func NewPostService(client v1.PostServiceClient) *PostService {
	return &PostService{client: client}
}

// Creating a new post.
func (s *PostService) CreatePost(ctx context.Context, input domain.CreatePostInput) (uuid.UUID, error) {
	// Getting author uuid from string.
	id, err := uuid.FromString(input.AuthorID)
	if err != nil {
		return uuid.Nil, err
	}

	// Create post.
	response, err := s.client.CreatePost(ctx, &v1.CreatePostRequest{
		AuthorId: id.Bytes(),
		Text:     input.Text,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.FromBytesOrNil(response.Id), nil
}

// Deleting a post.
func (s *PostService) DeletePost(ctx context.Context, id string) error {
	// Getting author uuid from string.
	authorID, err := uuid.FromString(id)
	if err != nil {
		return err
	}

	// Delete post.
	_, err = s.client.DeletePost(ctx, &v1.DeletePostRequest{Id: authorID.Bytes()})
	if err != nil {
		return err
	}

	return nil
}

// Updating a post.
func (s *PostService) UpdatePost(ctx context.Context, input domain.UpdatePostInput) error {
	// Getting post uuid from string.
	id, err := uuid.FromString(input.ID)
	if err != nil {
		return err
	}

	// Update post.
	_, err = s.client.UpdatePost(ctx, &v1.UpdatePostRequest{Id: id.Bytes(), Text: input.Text})
	if err != nil {
		return err
	}

	return nil
}

// Getting a post.
func (s *PostService) GetPost(ctx context.Context, id string) (*domain.Post, error) {
	// Getting post uuid from string.
	postID, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	// Getting post by id.
	post, err := s.client.GetPostById(ctx, &v1.GetPostByIdRequest{Id: postID.Bytes()})
	if err != nil {
		return nil, err
	}

	return &domain.Post{
		ID:        id,
		Author:    &domain.User{ID: uuid.FromBytesOrNil(post.AuthorId).String()},
		Text:      post.Text,
		CreatedAt: post.CreatedAt.AsTime(),
		UpdatedAt: post.UpdatedAt.AsOptionalTime(),
	}, nil
}
