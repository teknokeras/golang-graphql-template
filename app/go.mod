module github.com/teknokeras/golang-graphql-template/app

replace github.com/teknokeras/golang-graphql-template/app/schema => ./schema

replace github.com/teknokeras/golang-graphql-template/app/graphqlutils => ./graphqlutils

replace github.com/teknokeras/golang-graphql-template/app/modules => ./modules

replace github.com/teknokeras/golang-graphql-template/app/modules/core/role => ./modules/core/role

replace github.com/teknokeras/golang-graphql-template/app/modules/core/user => ./modules/core/user

require (
	github.com/go-pg/pg v7.1.7+incompatible
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	mellium.im/sasl v0.2.1 // indirect
)
