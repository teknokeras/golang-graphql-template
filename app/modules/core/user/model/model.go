package model

import (
	"errors"
	
	"github.com/teknokeras/golang-graphql-template/app/graphqlutils"
	"github.com/teknokeras/golang-graphql-template/app/db"
)

type User struct{
	Id 		  int64
	Name      string
	Email     string
	Password  string
	RoleId	  int64
}

type UserList struct {
	PageInfo 	graphqlutils.PageInfo
	TotalCount 	int 
	Users 		[]User
}

func GetUserById(id int) (User, error){
	u := User{Id: int64(id)}

	err := db.Engine.Select(&u)
	if err != nil {
		return User{}, errors.New("User Not Found")
    }
    return u, nil
}

func GetUserByEmail(email string) (User, error){
	r := User{}

	err := db.Engine.Model(&r).Where("email=?", email).Select()
	if err != nil {
		return User{}, errors.New("User Not Found")
    }
    return r, nil
}

func GetUsersPaginated(first int, after int) ([]User, error){
	var users []User
	err := db.Engine.Model(&users).Where("id>?", after).Limit(first).Order("id ASC").Select()

	if err != nil{
		return users, errors.New("Users Not Found")
	}

	if (len(users) == 0){
		return users, errors.New("No Users match the criteria are found")
	}

	return users, nil
}

func GetUsersByNamePaginated(first int, after int, name string) ([]User, error){
	var users []User
	err := db.Engine.Model(&users).Where("name LIKE ?", "%"+name+"%").Where("id>?", after).Limit(first).Order("id ASC").Select()

	if err != nil{
		return users, errors.New("Users Not Found")
	}

	if (len(users) == 0){
		return users, errors.New("No Users match the criteria are found")
	}

	return users, nil
}

func GetUsersByRolePaginated(first int, after int, roleId int) ([]User, error){
	var users []User
	err := db.Engine.Model(&users).Where("roleId=?", roleId).Where("id>?", after).Limit(first).Order("id ASC").Select()

	if err != nil{
		return users, errors.New("Users Not Found")
	}

	if (len(users) == 0){
		return users, errors.New("No Users match the criteria are found")
	}

	return users, nil
}

func GetPrevNextExistence(users []User) (bool, bool){
	user := users[0]
	prevCount, _ := db.Engine.Model((*User)(nil)).Where("id < ?", user.Id).Order("id ASC").Count()
	
	user = users[len(users) - 1]
	nextCount, _ := db.Engine.Model((*User)(nil)).Where("id > ?", user.Id).Order("id ASC").Count()

	return prevCount > 0, nextCount > 0
}
