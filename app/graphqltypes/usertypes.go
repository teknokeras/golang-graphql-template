package graphqltypes

import (
    "errors"

    "github.com/graphql-go/graphql"
    
    "github.com/teknokeras/golang-graphql-template/app/db"

    userModel "github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
    roleModel "github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UserInputType = graphql.NewInputObject(
    graphql.InputObjectConfig{
        Name: "UserInputType",
        Fields: graphql.InputObjectConfigFieldMap{
            "id": &graphql.InputObjectFieldConfig{
                Type: graphql.Int,
            },
            "name": &graphql.InputObjectFieldConfig{
                Type: graphql.String,
            },
            "email": &graphql.InputObjectFieldConfig{
                Type: graphql.String,
            },
            "password": &graphql.InputObjectFieldConfig{
                Type: graphql.String,
            },
            "roleId": &graphql.InputObjectFieldConfig{
                Type: graphql.Int,
            },
        },
    },
)

var roleField = &graphql.Field{
            Type: RoleType,
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {

                if user, ok := p.Source.(userModel.User); ok {

                    role := roleModel.Role{Id: user.RoleId}

                    err := db.Engine.Select(&role)
                    if err != nil {
                        return nil, errors.New("Cannot get User's Role")
                    }

                    return role, nil
                }else{
                    return nil, errors.New("Role is missing")
                }
            },
        }

var UserListType = graphql.NewObject(graphql.ObjectConfig{
        Name: "UserList",
        Fields: graphql.Fields{
            "users": &graphql.Field{
                Type:        graphql.NewList(UserType),
                Description: "Users of the list.",
            },
            "pageInfo": &graphql.Field{
                Type:        graphql.NewNonNull(PageInfoType),
                Description: "Information to aid in pagination.",
            },
            "totalCount": &graphql.Field{
                Type:        graphql.NewNonNull(graphql.Int),
                Description: "Count of all list users.",
            },
        },
    })

func init(){
    UserType.AddFieldConfig("role", roleField)  
}
