package dao

import (
	"database/sql"
	"fmt"
	"go-mysql/go/model"
)

func GetUser(userId int64) (model.User, error) {
	sqlQuery := "SELECT id, name, age, status FROM users WHERE id = ?"
	stmt, err := GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	var user model.User
	if err != nil {
		return user, err
	}
	res, err := stmt.Query(userId)
	defer closeRows(res)
	if err != nil {
		return user, err
	}
	if res.Next() {
		res.Scan(&user.UserID, &user.Name, &user.Age, &user.Status)
	}
	return user, err
}

func GetUsers() (model.Users, error) {
	sqlQuery := "SELECT id, name, age, status FROM users ORDER BY id DESC "
	stmt, err := GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)

	if err != nil {
		fmt.Println(err)
	}
	rows, err := stmt.Query()
	defer closeRows(rows)
	if err != nil {
		fmt.Println(err)
	}

	// Call the struct to be rendered on template
	n := model.User{}
	resUsers := model.Users{}
	var user model.User

	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Name, &user.Age, &user.Status)

		if err != nil {
			panic(err.Error())
		}

		// Get the Scan into the Struct
		n.UserID = user.UserID
		n.Name = user.Name
		n.Age = user.Age
		n.Status = user.Status

		resUsers = append(resUsers, n)
	}
	return resUsers, err
}

// AddUser method
func AddUser(user model.User) (int64, int64, error) {
	sqlQuery := "INSERT users SET name = ?, age = ?"
	stmt, err := GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(user.Name, user.Age)
	if err != nil {
		return 0, 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return rowsAffected, lastInsertedId, err
}

func UpdateUser(user model.User) (int64, error) {
	sqlQuery := "UPDATE users SET name = ?, age = ? WHERE id = ?"
	stmt, err := GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(user.Name, user.Age, user.UserID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

func DeleteUser(userId int64) (int64, error) {
	sqlQuery := "DELETE FROM users WHERE id = ?"
	stmt, err := GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(userId)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

func closeRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}

func closeStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}
