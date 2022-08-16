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
	// Post id.
	Id ksuid.KSUID `json:"id"`
	// Post author.
	Author *User `json:"author"`
	// Post text.
	Text string `json:"text"`
	// Post updated date.
	UpdatedAt *time.Time `json:"updatedAt"`
	// Post attachments.
	Attachments []string `json:"attachments"`
}

func (Post) IsNode() {}

// List of post owned by the subject.
type PostConnection struct {
	// A list of nodes.
	Nodes []*Post `json:"nodes"`
	// A list of edges.
	Edges []*PostEdge `json:"edges"`
	// Information to aid in pagination.
	PageInfo *PageInfo `json:"pageInfo"`
}

// An edge in a post connection.
type PostEdge struct {
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
	// The item at the end of the edge.
	Node *Post `json:"node"`
}

// Create post input.
type CreatePostInput struct {
	// Author id.
	AuthorId ksuid.KSUID
	// Post text.
	Text string `json:"text"`
	// Post attachments.
	Attachments []*UploadFile `json:"attachments"`
}

// Update post input.
type UpdatePostInput struct {
	// Post id.
	Id ksuid.KSUID `json:"id"`
	// Author id.
	AuthorId ksuid.KSUID
	// Post text.
	Text string `json:"text"`
}
