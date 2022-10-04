package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/coksnuss/entgql-noders-bug/ent"
	"github.com/coksnuss/entgql-noders-bug/ent/schema/types"
	"github.com/coksnuss/entgql-noders-bug/graph/generated"
	"github.com/google/uuid"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, post *ent.CreatePostInput) (*ent.Post, error) {
	record := r.Client.Post.Create().
		SetID(types.NewPostID(uuid.New())).
		SetTitle(post.Title)

	if post.CreatorID != nil {
		record.SetCreatorID(*post.CreatorID)
	}

	return record.Save(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
