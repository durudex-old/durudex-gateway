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

package domain

import "time"

// Post type.
type Post struct {
	ID          string     `json:"id"`
	AuthorID    string     `json:"authorId"`
	Text        string     `json:"text"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	Attachments []string   `json:"attachments"`
}

func (Post) IsNode() {}

// Create post input.
type CreatePostInput struct {
	AuthorID    string
	Text        string `json:"text"`
	Attachments []*UploadFile
}

// Update post input.
type UpdatePostInput struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
