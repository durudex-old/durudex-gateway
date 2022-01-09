/*
	Copyright © 2021-2022 Durudex

	This file is part of Durudex: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	Durudex is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with Durudex. If not, see <https://www.gnu.org/licenses/>.
*/

package graphql

import (
	"github.com/durudex/durudex-gateway/internal/delivery/graphql/generated"
	"github.com/durudex/durudex-gateway/internal/service"
)

type Resolver struct {
	service *service.Service
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

// Creating a new resolver.
func NewResolver(service *service.Service) *Resolver {
	return &Resolver{service: service}
}

// Mutation resolver.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query resolver.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
