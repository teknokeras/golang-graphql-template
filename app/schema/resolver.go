package schema

import (	
	graphql "github.com/graph-gophers/graphql-go"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/user"
	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
	"github.com/teknokeras/golang-graphql-template/app/db"
)	


// Resolver is the root resolver
type Resolver struct {
	db *db.DB
}

//allRoles(first: Int, offset: Int): RoleConnection
func (r *Resolver) AllRoles(ctx context.Context, args struct{ 
	first int
	offset int
}) (*role.RoleConnectionResolver, error) {
	rcr := role.RoleConnectionResolver{
		db *db.DB,
		first: first,
		offset: offset,
	}
	return &rcr, nil	
}

//getRoleById(id: ID!): Role
func (r *Resolver) GetRoleById(ctx context.Context, args struct{ ID graphql.ID }) (*role.RoleResolver, error) {
	roleID, err := graphqlutils.gqlIDToUint(args.ID)
	
	role, err := r.db.getRole(ctx, roleID)
	

	rcr := role.RoleResolver{
		first: first,
		offset: offset,
	}
	return &rcr, nil	
}

//getRoleByName(name: String!, first: Int, offset: Int): RoleConnection
func (r *Resolver) GetRoleByName(ctx context.Context, args struct{ 
	name string
	first int
	offset int
}) (*role.RoleConnectionResolver, error) {
		
}


//allUsers(first: Int, offset: Int): [UserConnection]
func (r *Resolver) AllUsers(ctx context.Context, args struct{ 
	first int
	offset int
}) (*user.UserConnectionResolver, error) {
	
}

//getUserById(id: ID!): User
func (r *Resolver) GetUserById(ctx context.Context, args struct{ ID graphql.ID }) (*user.UserResolver, error) {
	
}

//getUserByName(name: String!, first: Int, offset: Int): [UserConnection]
func (r *Resolver) GetUserByName(ctx context.Context, args struct{ 
	name string
	first int
	offset int
}) (*role.UserConnectionResolver, error) {
		
}

//getUserByEmail(email: String!): User
func (r *Resolver) GetUserByEmail(ctx context.Context, args struct{ email string }) (*role.UserResolver, error) {
		
}

//getUserByRoleName(role: String!, first: Int, offset: Int): [UserConnection]
func (r *Resolver) GetUserByRoleName(ctx context.Context, args struct{ 
	role string
	first int
	offset int
}) (*role.UserConnectionResolver, error) {
		
}
