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
	"encoding/base64"

	"github.com/segmentio/ksuid"
)

const (
	UserCtx   string = "userId"
	AuthorCtx string = "authorId"
	IpCtx     string = "Ip"
)

// Query sorting options.
type SortOptions struct {
	First  *int32
	Last   *int32
	Before ksuid.KSUID
	After  ksuid.KSUID
}

// Creating a new query sort options.
func NewSortOptions(first, last *int, before, after *string) (SortOptions, error) {
	sort := SortOptions{}

	// Check query first sort option.
	if first != nil {
		f := int32(*first)
		sort.First = &f
	}
	// Check query last sort option.
	if last != nil {
		l := int32(*last)
		sort.Last = &l
	}
	// Check query before sort option.
	if before != nil {
		b, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return SortOptions{}, err
		}

		sort.Before = ksuid.FromBytesOrNil(b)
	}
	// Check query after sort option.
	if after != nil {
		a, err := base64.StdEncoding.DecodeString(*after)
		if err != nil {
			return SortOptions{}, err
		}

		sort.After = ksuid.FromBytesOrNil(a)
	}

	return sort, nil
}
