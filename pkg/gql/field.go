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

// Getting selections fields.
func GetSelectionsFields(ctx context.Context, target string) []string {
	var fields []string

	// Getting selections.
	for _, i := range graphql.CollectFieldsCtx(ctx, nil) {
		// Check is target field.
		if i.Name == target {
			reqCtx := graphql.GetOperationContext(ctx)

			// Getting selections fields.
			for _, field := range graphql.CollectFields(reqCtx, i.Selections, nil) {
				fields = append(fields, field.Name)
			}
		}
	}

	return fields
}
