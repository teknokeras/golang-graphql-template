package role

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)	
type RoleEdge struct {
	cursor graphql.ID
	node   RoleResolver
}

// Cursor resolves the cursor for pagination
func (u *RoleEdge) Cursor(ctx context.Context) graphql.ID {
	return u.cursor
}

// Node resolves the node for pagination
func (u *RoleEdge) Node(ctx context.Context) *RoleResolver {
	return &u.node
}
