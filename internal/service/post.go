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

// Post service interface.
type Post interface {
	CreatePost(ctx context.Context, input domain.CreatePostInput) (string, error)
	DeletePost(ctx context.Context, id string) (bool, error)
	GetPost(ctx context.Context, id string) (domain.Post, error)
	UpdatePost(ctx context.Context, input domain.UpdatePostInput) (bool, error)
}

// Post service structure.
type PostService struct{}

// Creating a new post service.
func NewPostService() *PostService {
	return &PostService{}
}

// Creating a new post.
func (s *PostService) CreatePost(ctx context.Context, input domain.CreatePostInput) (string, error) {
	return "", nil
}

// Deleting a post.
func (s *PostService) DeletePost(ctx context.Context, id string) (bool, error) {
	return true, nil
}

// Getting a post.
func (s *PostService) GetPost(ctx context.Context, id string) (domain.Post, error) {
	return domain.Post{}, nil
}

// Updating a post.
func (s *PostService) UpdatePost(ctx context.Context, input domain.UpdatePostInput) (bool, error) {
	return true, nil
}
