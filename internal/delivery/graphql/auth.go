/*
	Copyright © 2021 Durudex

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
	"context"

	"github.com/Durudex/durudex-gateway/internal/delivery/graphql/generated"
	"github.com/Durudex/durudex-gateway/internal/delivery/graphql/model"
)

type mutationResolver struct{ *Resolver }

// Mutation resolver
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Sign Up resolver
func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpInput) (*model.SignUp, error) {
	return &model.SignUp{}, nil
}

// Sign In resolver
func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.SignIn, error) {
	return &model.SignIn{}, nil
}
