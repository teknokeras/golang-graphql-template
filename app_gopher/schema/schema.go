package schema

import (
	"os"
	"io/ioutil"
	"strings"
)

var pwd, _ = os.Getwd()
var modulePath = pwd+"/modules"
var moduleSchemaPath = "schema/"
var schemaTypeFilename = "type.graphql"
var schemaQueryFilename = "query.graphql"
var schemaMutationFilename = "mutation.graphql"

func readSchemaFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func wrapModuleSchemaTypePath(moduleName string) string{
	var completePath strings.Builder

	completePath.WriteString(modulePath)
	completePath.WriteString(moduleName)
	completePath.WriteString(moduleSchemaPath)
	completePath.WriteString(schemaTypeFilename)

	return completePath.String()
}

func wrapModuleSchemaQueryPath(moduleName string) string{
	var completePath strings.Builder

	completePath.WriteString(modulePath)
	completePath.WriteString(moduleName)
	completePath.WriteString(moduleSchemaPath)
	completePath.WriteString(schemaQueryFilename)

	return completePath.String()
}

func wrapModuleSchemaMutationPath(moduleName string) string{
	var completePath strings.Builder

	completePath.WriteString(modulePath)
	completePath.WriteString(moduleName)
	completePath.WriteString(moduleSchemaPath)
	completePath.WriteString(schemaMutationFilename)

	return completePath.String()
}

func buildSchemaQuery(moduleName string) string {
	var queryFilePath = wrapModuleSchemaQueryPath(moduleName)
	var queryString, err = readSchemaFile(queryFilePath)

	if err != nil{
		return ""
	}

	return queryString
}

func buildSchemaMutation(moduleName string) string{
	var mutationFilePath = wrapModuleSchemaMutationPath(moduleName)
	var mutationString, err = readSchemaFile(mutationFilePath)

	if err != nil{
		return ""
	}

	return mutationString
}

func buildSchemaType(moduleName string) string{
	var typePath = wrapModuleSchemaTypePath(moduleName)
	var typeString, err = readSchemaFile(typePath)

	if err != nil{
		return ""
	}

	return typeString
}

func Build() string{
	var schemaBuilder strings.Builder

	var schemaHead = `
	schema {
	  query: Query
	  mutation: Mutation
	}
	`

	schemaBuilder.WriteString(schemaHead)

	//Read core schema (user and role)
	coreRolePath := "/core/role/"
	coreUserPath := "/core/user/"

	schemaBuilder.WriteString(`
	type Query {
	`)

	schemaBuilder.WriteString(buildSchemaQuery(coreRolePath))
	schemaBuilder.WriteString(buildSchemaQuery(coreUserPath))

	schemaBuilder.WriteString(`
	}
	`)

	schemaBuilder.WriteString(`
	type Mutation {
	`)

	schemaBuilder.WriteString(buildSchemaMutation(coreRolePath))
	schemaBuilder.WriteString(buildSchemaMutation(coreUserPath))

	schemaBuilder.WriteString(`
	}
	
	`)

	schemaBuilder.WriteString(buildSchemaType(coreRolePath))
	schemaBuilder.WriteString(`
		`)
	schemaBuilder.WriteString(buildSchemaType(coreUserPath))

	schemaBuilder.WriteString(`

	type PageInfo {
      startCursor: ID
      endCursor: ID
      hasNextPage: Boolean!
      hasPreviousPage: Boolean!
    }
	`)

	return schemaBuilder.String()
}