package resolver

import 	(
	"context"

    "github.com/go-pg/pg"
	graphql "github.com/graph-gophers/graphql-go"

	user "github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
)

type UserResolver struct {
	db     *pg.DB
	user   user.User
}

func (r *UserResolver) ID(ctx context.Context) *graphql.ID {
	return graphqlutils.gqlIDP(r.role.Id)
}

func (r *RoleResolver) Name(ctx context.Context)  string {
	return r.role.Name
}