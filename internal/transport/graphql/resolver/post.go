package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"

	"github.com/durudex/durudex-gateway/internal/domain"
	"github.com/durudex/durudex-gateway/internal/transport/graphql/generated"
	"github.com/durudex/durudex-gateway/pkg/gql"
	"github.com/segmentio/ksuid"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input domain.CreatePostInput) (ksuid.KSUID, error) {
	var err error

	// Parsing author id.
	input.AuthorId, err = ksuid.Parse(ctx.Value(domain.UserCtx).(string))
	if err != nil {
		return ksuid.Nil, err
	}

	// Create post.
	id, err := r.service.Post.Create(ctx, input)
	if err != nil {
		return ksuid.Nil, err
	}

	return id, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id ksuid.KSUID) (bool, error) {
	// Parsing post id.
	postId, err := ksuid.Parse(ctx.Value(domain.UserCtx).(string))
	if err != nil {
		return false, err
	}

	// Delete post.
	if err := r.service.Post.Delete(ctx, id, postId); err != nil {
		return false, err
	}

	return true, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input domain.UpdatePostInput) (bool, error) {
	var err error

	// Parsing author id from string.
	input.AuthorId, err = ksuid.Parse(ctx.Value(domain.UserCtx).(string))
	if err != nil {
		return false, err
	}

	// Update post.
	if err := r.service.Post.Update(ctx, input); err != nil {
		return false, err
	}

	return true, nil
}

// Edges is the resolver for the edges field.
func (r *postConnectionResolver) Edges(ctx context.Context, obj *domain.PostConnection) ([]*domain.PostEdge, error) {
	edges := make([]*domain.PostEdge, len(obj.Nodes))

	for i, node := range obj.Nodes {
		edges[i] = &domain.PostEdge{
			Cursor: base64.StdEncoding.EncodeToString(node.Id.Bytes()),
			Node:   node,
		}
	}

	return edges, nil
}

// PageInfo is the resolver for the pageInfo field.
func (r *postConnectionResolver) PageInfo(ctx context.Context, obj *domain.PostConnection) (*domain.PageInfo, error) {
	start := base64.StdEncoding.EncodeToString(obj.Nodes[0].Id.Bytes())
	end := base64.StdEncoding.EncodeToString(obj.Nodes[len(obj.Nodes)-1].Id.Bytes())

	return &domain.PageInfo{
		StartCursor: &start,
		EndCursor:   &end,
	}, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id ksuid.KSUID) (*domain.Post, error) {
	// Getting a post.
	post, err := r.service.Post.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Getting author selections fields.
	fields := gql.GetSelectionsFields(ctx, "author")

	// Check author selections fields.
	if !(len(fields) == 1 && fields[0] == "id") {
		// Getting post author.
		user, err := r.service.User.GetById(ctx, post.Author.Id)
		if err != nil {
			return nil, err
		}

		post.Author = user
	}

	return post, nil
}

// PostConnection returns generated.PostConnectionResolver implementation.
func (r *Resolver) PostConnection() generated.PostConnectionResolver {
	return &postConnectionResolver{r}
}

type postConnectionResolver struct{ *Resolver }
