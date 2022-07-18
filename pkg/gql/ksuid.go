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
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/segmentio/ksuid"
)

// Marshal a KSUID to a string.
func MarshalKSUID(v ksuid.KSUID) graphql.Marshaler {
	return graphql.MarshalString(v.String())
}

// Unmarshal a KSUID from a string.
func UnmarshalKSUID(v any) (ksuid.KSUID, error) {
	switch v := v.(type) {
	case string:
		return ksuid.Parse(v)
	default:
		return ksuid.Nil, fmt.Errorf("ksuid must be a string")
	}
}
