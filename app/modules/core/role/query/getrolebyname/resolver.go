package getrolebyname

import (
	"errors"
	
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/db"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

func Resolve(params graphql.ResolveParams, database *db.Database) (interface{}, error){

	if (params.Args["name"] == nil){
		return nil, errors.New("role input argument is mandatory")
	}

	name, _ := params.Args["name"].(string)

	r := model.Role{Name: name}

	err := database.Engine.Select(&r)
	if err == nil {
		return nil, errors.New("Cannot get role from DB")
    }

	return r, nil

}