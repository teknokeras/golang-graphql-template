package roleconn

import (
	"context"
	
	"github.com/teknokeras/golang-graphql-template/app/schema"
	"github.com/teknokeras/golang-graphql-template/app/db"
	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
	role "github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
	roleResolver "github.com/teknokeras/golang-graphql-template/app/modules/core/role/resolver/role"
)

type RoleConnectionResolver struct {
	Db 			*db.Database
	TotalRoles 	int
	Roles 		[]role.Role
}

func (r *RoleConnectionResolver) TotalCount(ctx context.Context) int{
	return TotalRoles
}

func (r *RoleConnectionResolver) Edges(ctx context.Context) (*[]*roleResolver.RoleEdge, error){

	l := make([]*roleResolver.RoleEdge, len(r.Roles))

	for i := range l {
		Id := graphqlutils.encodeCursor(r.Roles[i].Id)
		l[i] = &roleResolver.RoleEdge{
			cursor: Id,
			node: roleResolver.RoleResolver{
				Db: r.Db,
				M:  r.Roles[i],
				Id: Id,
			},
		}
	}

	return &l, nil
}

func (r *RoleConnectionResolver) PageInfo(ctx context.Context) (*schema.PageInfo, error) {

	var roles []role.Role

	err := r.Db.Model(&roles).Order("Id ASC").Limit(1).Select()

    if err != nil {
    	return nil, errors.New("Cannot get roles from DB")
    }

    prevPageExists := roles[0].Id < r.Roles[0].Id 

	err = r.Db.Model(&roles).Order("Id DESC").Limit(1).Select()

    if err != nil {
    	return nil, errors.New("Cannot get roles from DB")
    }

    nextPageExists := roles[len(roles) - 1].Id > r.Roles[len(r.Roles) - 1].Id 

	p := PageInfo{
		startCursor:     graphqlutils.EncodeCursor(r.Roles[0].Id),
		endCursor:       graphqlutils.EncodeCursor(r.Roles[len(r.Roles) - 1].Id),
		hasNextPage:     nextPageExists,
		hasPreviousPage: prevPageExists,
	}
	return &p, nil
}