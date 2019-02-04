package model

import (
	"errors"

	"github.com/teknokeras/golang-graphql-template/app/db"
	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
	user "github.com/teknokeras/golang-graphql-template/app/modules/core/user/model"
)

type Role struct {
	Id    int64
	Name  string
	Users []*user.User
}

type RoleList struct {
	PageInfo   graphqlutils.PageInfo
	TotalCount int
	Roles      []Role
}

func GetRoleById(id int) (Role, error) {
	r := Role{Id: int64(id)}

	err := db.Engine.Select(&r)
	if err != nil {
		return Role{}, errors.New("Role Not Found")
	}
	return r, nil
}

func GetRoleByName(name string) (Role, error) {
	r := Role{}

	err := db.Engine.Model(&r).Where("name=?", name).Select()
	if err != nil {
		return Role{}, errors.New("Role Not Found")
	}
	return r, nil
}

func GetRolesPaginated(first int, after int) ([]Role, error) {
	var roles []Role
	err := db.Engine.Model(&roles).Where("id > ?", after).Order("id ASC").Limit(first).Select()

	if err != nil {
		return roles, errors.New("Roles Not Found")
	}

	if len(roles) == 0 {
		return roles, errors.New("No Roles match the criteria are found")
	}

	return roles, nil
}

func GetPrevNextExistence(after int, roles []Role) (bool, bool) {
	prevCount, _ := db.Engine.Model((*Role)(nil)).Where("id < ?", after).Order("id ASC").Count()

	nextCount, _ := db.Engine.Model((*Role)(nil)).Where("id > ?", roles[len(roles)-1].Id).Order("id ASC").Count()

	return prevCount > 0, nextCount > 0
}

func CreateRole(name string) (Role, error) {

	if r, err := GetRoleByName(name); err == nil {
		//role exist then reject this request
		return Role{}, errors.New("This role already exists")
	} else {
		r = Role{Name: name}
		if err = db.Engine.Insert(&r); err != nil {
			return Role{}, err
		}

		return r, nil
	}
}

func UpdateRole(role Role) (Role, error) {
	r := Role{Id: role.Id}

	if err := db.Engine.Select(&r); err != nil {
		return Role{}, err
	} else {
		r.Name = role.Name

		if err = db.Engine.Update(&r); err != nil {
			return Role{}, err
		} else {
			return r, nil
		}
	}
}

func DeleteRole(id int) error {
	r := Role{Id: int64(id)}

	if err := db.Engine.Select(&r); err != nil {
		return errors.New("Role Does not Exists")
	} else {
		if err = db.Engine.Delete(&r); err != nil {
			return errors.New("Role Cannot be deleted")
		} else {
			return nil
		}
	}
}
