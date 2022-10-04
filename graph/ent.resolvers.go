package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"regexp"

	"entgo.io/contrib/entgql"
	"github.com/coksnuss/entgql-noders-bug/ent"
	"github.com/coksnuss/entgql-noders-bug/graph/generated"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	prefixMatcher := regexp.MustCompile("^/([^/]+)/")
	resolver := ent.WithNodeType(func(_ context.Context, id string) (string, error) {
		matches := prefixMatcher.FindStringSubmatch(id)
		if matches == nil {
			return "", entgql.ErrNodeNotFound(id)
		}
		return matches[1], nil
	})

	return r.Client.Noder(ctx, id, resolver)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	prefixMatcher := regexp.MustCompile("^/([^/]+)/")
	resolver := ent.WithNodeType(func(_ context.Context, id string) (string, error) {
		matches := prefixMatcher.FindStringSubmatch(id)
		if matches == nil {
			return "", entgql.ErrNodeNotFound(id)
		}
		return matches[1], nil
	})

	return r.Client.Noders(ctx, ids, resolver)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.PostConnection, error) {
	return r.Client.Post.Query().Paginate(ctx, after, first, before, last)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.UserConnection, error) {
	return r.Client.User.Query().Paginate(ctx, after, first, before, last)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
