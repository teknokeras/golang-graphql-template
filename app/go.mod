module github.com/teknokeras/golang-graphql-template/app

replace github.com/teknokeras/golang-graphql-template/app/schema => ./schema

replace github.com/teknokeras/golang-graphql-template/app/modules/core/role => ./modules/core/role

replace github.com/teknokeras/golang-graphql-template/app/modules/core/user => ./modules/core/user

require (
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/lib/pq v1.0.0 // indirect
)
