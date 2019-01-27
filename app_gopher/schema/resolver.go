package schema

import (
	"context"
	
	graphql "github.com/graph-gophers/graphql-go"

	"github.com/teknokeras/golang-graphql-template/app/db"	
)

// Resolver is the root resolver
type Resolver struct {
	Db *db.Database
}

type PageInfo struct {
	startCursor     graphql.ID
	endCursor       graphql.ID
	hasNextPage     bool
	hasPreviousPage bool
}

// StartCursor ...
func (u *PageInfo) StartCursor(ctx context.Context) *graphql.ID {
	return &u.startCursor
}

// EndCursor ...
func (u *PageInfo) EndCursor(ctx context.Context) *graphql.ID {
	return &u.endCursor
}

// HasNextPage returns true if there are more results to show
func (u *PageInfo) HasNextPage(ctx context.Context) bool {
	return u.hasNextPage
}

// HasPreviousPage returns true if there are results behind the current cursor position
func (u *PageInfo) HasPreviousPage(ctx context.Context) bool {
	return u.hasPreviousPage
}
