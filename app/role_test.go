package main

import (
	"testing"

	"github.com/graphql-go/graphql"

	appSchema "github.com/teknokeras/golang-graphql-template/app/schema"
)

func TestGetRoles(t *testing.T) {

	query := `
		{
			GetRoles(first: 10,after: 0){
			    totalCount
			    pageInfo{
			      hasNextPage
			      hasPreviousPage
			    }
			    roles{
			      id
			      name
			    }
			}
		}
	`

	params := graphql.Params{Schema: appSchema.Schema, RequestString: query}

	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		t.Error(r.Errors[0])
	}
}

func TestGetRoleById(t *testing.T) {

	query := `
		{
			GetRoleById(id: 1){
				id 
				name 
				users {
					id 
					name
				}
			}
		}
	`

	params := graphql.Params{Schema: appSchema.Schema, RequestString: query}

	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		t.Error(r.Errors[0])
	}
}

func TestGetRoleByName(t *testing.T) {

	query := `
		{
			GetRoleByName(name: "ADMINISTRATOR"){
				id 
				name 
				users {
					id 
					name
				}
			}
		}
	`

	params := graphql.Params{Schema: appSchema.Schema, RequestString: query}

	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		t.Error(r.Errors[0])
	}
}
