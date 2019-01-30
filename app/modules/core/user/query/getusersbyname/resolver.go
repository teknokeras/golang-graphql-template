package getusersbyname

import (
	"github.com/teknokeras/golang-graphql-template/app/db"
	"github.com/graphql-go/graphql"
)

func Resolve(params graphql.ResolveParams, database db.Database) (interface{}, error){
	return nil, nil
}
/*
import (
	"errors"
	
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/core/user/args"
)


func Resolve(params graphql.ResolveParams) (interface{}, error){
	 // marshall and cast the argument value
	text, _ := params.Args["name"].(string)

	// figure out new id
	newID := RandStringRunes(8)

	// perform mutation operation here
	// for e.g. create a Todo and save to DB.
	newTodo := Todo{
		ID:   newID,
		Text: text,
		Done: false,
	}

	TodoList = append(TodoList, newTodo)

	// return the new Todo object that we supposedly save to DB
	// Note here that
	// - we are returning a `Todo` struct instance here
	// - we previously specified the return Type to be `todoType`
	// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
	return newTodo, nil
}*/