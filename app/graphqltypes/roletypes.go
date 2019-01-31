package graphqltypes

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
	
	"github.com/teknokeras/golang-graphql-template/app/db"

	userModel "github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
	roleModel "github.com/teknokeras/golang-graphql-template/app/modules/core/role/model"
)

var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
   	},
})

var RoleInputType = graphql.NewInputObject(
    graphql.InputObjectConfig{
        Name: "RoleInputType",
        Fields: graphql.InputObjectConfigFieldMap{
            "id": &graphql.InputObjectFieldConfig{
                Type: graphql.Int,
            },
            "name": &graphql.InputObjectFieldConfig{
                Type: graphql.String,
            },
        },
    },
)

var usersField = &graphql.Field{
		Type: graphql.NewList(UserType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if role, ok := p.Source.(roleModel.Role); ok {

				var users []userModel.User
				err := db.Engine.Model(&users).Where("role_id=?", role.Id).Limit(10).Select(&users)
				if err != nil {
					fmt.Println(err)
			        return nil, errors.New("No Users for this Role")
			    }

				return users, nil
			}else{
				return nil, errors.New("Role is missing")
			}
        },
	}

var RoleListType = graphql.NewObject(graphql.ObjectConfig{
		Name: "RoleList",
		Fields: graphql.Fields{
			"totalCount": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "Count of all list roles.",
			},
			"pageInfo": &graphql.Field{
				Type:        graphql.NewNonNull(PageInfoType),
				Description: "Information to aid in pagination.",
			},
			"roles": &graphql.Field{
				Type:        graphql.NewList(RoleType),
				Description: "Roles of the list.",
			},
		},
	})

func init() {
	RoleType.AddFieldConfig("users", usersField)	
}


