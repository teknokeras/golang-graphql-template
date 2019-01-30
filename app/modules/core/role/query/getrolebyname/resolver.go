package getrolebyname

import (
	"errors"

	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/db"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

func Resolve(params graphql.ResolveParams) (interface{}, error){

	if (params.Args["name"] == nil){
		return nil, errors.New("role input argument is mandatory")
	}

	name, _ := params.Args["name"].(string)

	r := new(model.Role)

	err := db.Engine.Model(r).Where("name=?", name).Select()
	if err != nil {
		return nil, errors.New("Cannot get role from DB")
    }

	return r, nil

}