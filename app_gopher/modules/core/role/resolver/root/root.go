package root

import (
	"context"
	"encoding/base64"
	"errors"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
	"github.com/teknokeras/golang-graphql-template/app/schema"
	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/resolver/roleconn"

	role "github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

type roleConnArgs struct {
	First *int32
	After *graphql.ID
}

//allRoles(first: Int, offset: Int): RoleConnection
func (r *schema.Resolver) AllRoles(ctx context.Context, args roleConnArgs) (*roleconn.RoleConnectionResolver, error) {

	var roles []role.Role

	from := 0
	limit := 10
	if args.After == nil{
		err := r.Db.Model(&roles).Order("Id ASC").Limit(1).Select()
		if err != nil {
			return nil, errors.New("Cannot access Role in DB")
		}
		from := roles[0].Id
	}else{
		from, err := graphqlutils.DecodeCursor(string(*args.After))
		if err != nil {
			return nil, errors.New("Cannot decode After ID")
		}
	}

	if args.First != nil {
		limit = int(*args.First)
	}

	err = r.Db.Model(&roles).Order("Id ASC"),Where("Id > ?", from).Limit(limit).Select()
    if err != nil {
    	return nil, errors.New("Cannot get roles from DB")
    }

    count, err := r.Db.Model((*role.Role)(nil)).Count()

    if err != nil {
    	return nil, errors.New("Cannot get role count from DB")
    }

    rcr := roleconn.RoleConnectionResolver{
		Db: r.Db,
		TotalRoles: count
		Roles: roles
	}
	
	return &rcr, nil
}
