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

package domain

import (
	"time"

	"github.com/segmentio/ksuid"
)

// Post type.
type Post struct {
	Id          ksuid.KSUID `json:"id"`
	Author      *User       `json:"author"`
	Text        string      `json:"text"`
	UpdatedAt   *time.Time  `json:"updatedAt"`
	Attachments []string    `json:"attachments"`
}

// List of post owned by the subject.
type PostConnection struct {
	Nodes []*Post `json:"nodes"`
}

func (Post) IsNode() {}

// Create post input.
type CreatePostInput struct {
	AuthorId    ksuid.KSUID
	Text        string `json:"text"`
	Attachments []*UploadFile
}

// Update post input.
type UpdatePostInput struct {
	Id       ksuid.KSUID `json:"id"`
	AuthorId ksuid.KSUID
	Text     string `json:"text"`
}
