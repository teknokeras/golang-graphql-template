package resolver

import (
	"errors"
	
	"github.com/graphql-go/graphql"

	"github.com/teknokeras/golang-graphql-template/app/modules/code/role/query/getroles/args"
	"github.com/teknokeras/golang-graphql-template/app/modules/code/role/query/getroles/resolver"
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
}