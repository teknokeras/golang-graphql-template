package resolver

import (
	"errors"
	
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/db"

	"github.com/teknokeras/golang-graphql-template/app/modules/code/role/model"
)

func Resolve(params graphql.ResolveParams, database db.Database) (interface{}, error){

	if (params.Args["id"] == nil){
		return nil, errors.New("'id' argument is mandatory")
	}

	id, _ := params.Args["id"].(int)

	r := model.Role{Id: id}

	err := db.Select(&r)
	if err == nil {
		return model.Role{}, nil
	}

	return r, nil
}