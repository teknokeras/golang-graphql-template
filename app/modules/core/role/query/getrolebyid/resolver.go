package getrolebyid

import (
	"errors"
	
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/db"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

func Resolve(params graphql.ResolveParams, database *db.Database) (interface{}, error){
	if (params.Args["id"] == nil){
		return nil, errors.New("'id' argument is mandatory")
	}

	id, _ := params.Args["id"].(int)

	r := model.Role{Id: int64(id)}

	err := database.Engine.Select(&r)
	if err == nil {
		return nil, errors.New("Role Not Found")
    }

	return r, nil
}