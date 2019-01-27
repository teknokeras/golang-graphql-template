package role

import (
	"context"
	
	graphql "github.com/graph-gophers/graphql-go"

	"github.com/teknokeras/golang-graphql-template/app/db"
	role "github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)


type RoleResolver struct {
	Db 	*db.Database
	M   role.Role
	Id 	graphql.ID
}

func (r *RoleResolver) Id(ctx context.Context) graphql.ID {
	return (r.Id)
}

func (r *RoleResolver) Name(ctx context.Context) string {
	return (r.M.Name)
}

Untuk users dalam role ini dibatasi saja sampe 10 misalnya. selanjutnya kalo client ingin lebih bisa pake query di user (getusers by role id)