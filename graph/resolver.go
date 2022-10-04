//go:generate go run github.com/99designs/gqlgen generate

package graph

import "github.com/coksnuss/entgql-noders-bug/ent"

type Resolver struct {
	Client *ent.Client
}
