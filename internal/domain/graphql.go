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

import "github.com/99designs/gqlgen/graphql"

// GraphQL error status codes.
const (
	CodeServerError         string = "SERVER_ERROR"
	CodeInternalServerError string = "INTERNAL_SERVER_ERROR"
	CodeInvalidArgument     string = "INVALID_ARGUMENT"
	CodeNotFound            string = "NOT_FOUND"
	CodeAlreadyExists       string = "ALREADY_EXISTS"
	CodeUnauthorized        string = "UNAUTHORIZED"
)

// GraphQL Node interface.
type Node interface {
	IsNode()
}

// Information about pagination in a connection.
type PageInfo struct {
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor"`
}

// Upload files input.
type UploadFile struct {
	// File id.
	Id int `json:"id"`
	// File data.
	File graphql.Upload `json:"file"`
}
