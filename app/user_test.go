package main

import (
	"testing"

	"github.com/graphql-go/graphql"

	appSchema "github.com/teknokeras/golang-graphql-template/app/schema"
)

func TestGetUsers(t *testing.T) {

	query := `
		{
			GetUsers(first: 10,after: 0){
			    totalCount
			    pageInfo{
			      hasNextPage
			      hasPreviousPage
			    }
			    users{
			      id
			      name
			      email
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

func TestGetUserById(t *testing.T) {

	query := `
		{
			GetUserById(id: 1){
				id 
				email 
				name 
				role {
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

func TestGetUsersByName(t *testing.T) {

	query := `
		{
			GetUsersByName(name: "root"){
				pageInfo {
					hasNextPage 
					hasPreviousPage
				} 
				totalCount 
				users{
					id 
					name 
					email
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

func TestGetUserByRole(t *testing.T) {

	query := `
		{
			GetUsersByRole(roleId: 1){
				pageInfo {
					hasNextPage 
					hasPreviousPage
				} 
				totalCount 
				users{
					id 
					name 
					email
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
