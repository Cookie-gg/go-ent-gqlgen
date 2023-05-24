package resolver

import (
	"go-ent-gqlgen/ent"

	"go-ent-gqlgen/internal"
	"github.com/99designs/gqlgen/graphql"
)

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return internal.NewExecutableSchema(internal.Config{
		Resolvers: &Resolver{client},
	})
}