package service

import (
	"learn-go/src/dao"
	"learn-go/src/model"
)

// AddUser method
func AddUser(name string, email string, age int8) (model.User, error) {
	user := model.User{Name: name, Email: email, Age: age}
	rowsAffected, lastInsertedId, err := dao.AddUser(user)
	if err == nil && rowsAffected > 0 {
		user.id = lastInsertedId
	}
	return user, err
}

// func UpdateStudent(id int64, name string, age int8) (model.Student, error) {
// 	student := model.Student{StudentID:id, Name:name, Age:age}
// 	rowsAffected, err := dao.UpdateStudent(student)
// 	if err == nil && rowsAffected == 0 {
// 		err = errors.New("No Data Found")
// 	}
// 	return student, err
// }

// func DeleteStudent(studentID int64) (bool, error) {
// 	rowsAffected, err := dao.DeleteStudent(studentID)
// 	if err == nil && rowsAffected == 0 {
// 		err = errors.New("No Data Found")
// 	}
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, err
// }

// func GetStudent(studentID int64) (model.Student, error) {
// 	return dao.GetStudent(studentID)
// }
