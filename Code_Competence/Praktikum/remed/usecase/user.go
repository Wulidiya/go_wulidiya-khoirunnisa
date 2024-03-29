package usecase

import (
	"fmt"
	"remed/middleware"
	"remed/model"
	"remed/model/payload"
	"remed/repository/database"
)

func LoginUser(user *model.User) (err error) {
	// check to db email and password
	err = database.LoginUser(user)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	// generate jwt
	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		fmt.Println("GetUser: Error Generate token")
		return
	}
	user.Token = token
	return
}

func CreateUser(req *payload.CreateUserRequest) (resp payload.CreateUserResponse, err error) {
	newUser := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err = database.CreateUser(newUser)
	if err != nil {
		return
	}
	// generate jwt
	token, err := middleware.CreateToken(int(newUser.ID))
	if err != nil {
		fmt.Println("GetUser: Error Generate token")
		return
	}
	newUser.Token = token
	err = database.UpdateUser(newUser)
	if err != nil {
		fmt.Println("UpdateUser: Error Update user")
		return
	}
	resp = payload.CreateUserResponse{
		UserID: newUser.ID,
		Token:  newUser.Token,
	}
	return
}

func GetUser(id uint) (user model.User, err error) {
	user.ID = id
	err = database.GetUser(&user)
	if err != nil {
		fmt.Println("GetUser: Error getting user from database")
		return
	}
	return
}

func GetListUsers() (users []model.User, err error) {
	users, err = database.GetUsers()
	if err != nil {
		fmt.Println("GetListUsers: Error getting users from database")
		return
	}
	return
}

func UpdateUser(user *model.User) (err error) {
	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser : Error updating user, err: ", err)
		return
	}

	return
}

func DeleteUser(id uint) (err error) {
	user := model.User{}
	user.ID = id
	err = database.DeleteUser(&user)
	if err != nil {
		fmt.Println("DeleteUser : error deleting user, err: ", err)
		return
	}

	return
}
