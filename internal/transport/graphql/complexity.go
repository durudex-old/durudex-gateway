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

package graphql

import "github.com/durudex/durudex-gateway/internal/transport/graphql/generated"

// Setting the complexity of the query.
func setComplexity(root *generated.ComplexityRoot) {
	root.User.Posts = filterComplexity
	root.Post.Author = doubleComplexity
}

// Filter the complexity of the query.
func filterComplexity(childComplexity int, first *int, last *int, before *string, after *string) int {
	switch {
	case first != nil:
		return childComplexity * *first
	case last != nil:
		return childComplexity * *last
	default:
		return 0
	}
}

// Double the complexity of the query.
func doubleComplexity(childComplexity int) int {
	return childComplexity * 2
}
