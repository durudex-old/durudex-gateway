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

package gql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// Check is field single.
func IsFieldSingle(ctx context.Context, target string) bool {
	var res bool

	for i, field := range graphql.CollectFieldsCtx(ctx, nil) {
		if field.Name == target {
			res = true
		}

		if i == 1 {
			return false
		}
	}

	return res
}
