package main

import (
	"testing"

	"github.com/graphql-go/graphql"

	appSchema "github.com/teknokeras/golang-graphql-template/app/schema"
)

func TestLoginPass(t *testing.T) {

	query := `
		{
			Login(email:"root@flask.com", password:"flaskiscool")
		}
	`

	params := graphql.Params{Schema: appSchema.Schema, RequestString: query}

	r := graphql.Do(params)

	if len(r.Errors) > 0 {
		t.Error("This should pass and no error but error occured during login pass test")
	}
}

func TestLoginFail(t *testing.T) {

	query := `
		{
			Login(email:"root@dsf.com", password:"fdf")
		}
	`

	params := graphql.Params{Schema: appSchema.Schema, RequestString: query}

	r := graphql.Do(params)

	if len(r.Errors) == 0 {
		t.Error("This should fail but no error occured during login fail test")
	}
}
