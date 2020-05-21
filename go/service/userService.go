package service

import (
	"errors"
	"go-mysql/go/dao"
	"go-mysql/go/model"
)

func GetUsers() (model.Users, error) {
	return dao.GetUsers()
}

func GetUser(userId int64) (model.User, error) {
	return dao.GetUser(userId)
}

// AddUser method
func AddUser(name string, age int8) (model.User, error) {
	user := model.User{Name: name, Age: age}
	rowsAffected, lastInsertedId, err := dao.AddUser(user)
	if err == nil && rowsAffected > 0 {
		user.UserID = lastInsertedId
	}
	return user, err
}

func UpdateUser(id int64, name string, age int8, status string) (model.User, error) {
	user := model.User{UserID: id, Name: name, Age: age, Status: status}
	rowsAffected, err := dao.UpdateUser(user)
	if err == nil && rowsAffected == 0 {
		err = errors.New("No Data Found")
	}
	return user, err
}

func DeleteUser(userId int64) (bool, error) {
	rowsAffected, err := dao.DeleteUser(userId)
	if err == nil && rowsAffected == 0 {
		err = errors.New("No Data Found")
	}
	if err != nil {
		return false, err
	}
	return true, err
}
