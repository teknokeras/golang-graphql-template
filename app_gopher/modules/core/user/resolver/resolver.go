package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"

	"github.com/teknokeras/golang-graphql-template/app/schema"
)

//allUsers(first: Int, offset: Int): [UserConnection]
func (r *schema.Resolver) AllUsers(ctx context.Context, args struct{ 
	first int
	offset int
}) (*user.UserConnectionResolver, error) {
	
}

//getUserById(id: ID!): User
func (r *schema.Resolver) GetUserById(ctx context.Context, args struct{ ID graphql.ID }) (*user.UserResolver, error) {
	
}

//getUserByName(name: String!, first: Int, offset: Int): [UserConnection]
func (r *schema.Resolver) GetUserByName(ctx context.Context, args struct{ 
	name string
	first int
	offset int
}) (*role.UserConnectionResolver, error) {
		
}

//getUserByEmail(email: String!): User
func (r *schema.Resolver) GetUserByEmail(ctx context.Context, args struct{ email string }) (*role.UserResolver, error) {
		
}

//getUserByRoleName(role: String!, first: Int, offset: Int): [UserConnection]
func (r *schema.Resolver) GetUserByRoleName(ctx context.Context, args struct{ 
	role string
	first int
	offset int
}) (*role.UserConnectionResolver, error) {
		
}
