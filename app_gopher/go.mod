module github.com/teknokeras/golang-graphql-template/app

replace github.com/teknokeras/golang-graphql-template/app/schema => ./schema

replace github.com/teknokeras/golang-graphql-template/app/graphqlutils => ./graphqlutils

replace github.com/teknokeras/golang-graphql-template/app/modules => ./modules

replace github.com/teknokeras/golang-graphql-template/app/modules/core/role => ./modules/core/role

replace github.com/teknokeras/golang-graphql-template/app/modules/core/user => ./modules/core/user

require (
	github.com/go-pg/pg v7.1.7+incompatible
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	golang.org/x/crypto v0.0.0-20180910181607-0e37d006457b
	golang.org/x/sys v0.0.0-20190123074212-c6b37f3e9285 // indirect
	mellium.im/sasl v0.2.1 // indirect
)
