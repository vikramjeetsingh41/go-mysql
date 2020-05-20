package dao

import (
	// go
	. "database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *DB

// InitializeMySQL Initialize database connection
func InitializeMySQL() {
	dBConnection, err := Open("mysql", "root:@(localhost:3306)/golang")
	if err != nil {
		fmt.Println("Connection Failed!!")
	}
	err = dBConnection.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	}
	db = dBConnection
	dBConnection.SetMaxOpenConns(10)
	dBConnection.SetMaxIdleConns(5)
	dBConnection.SetConnMaxLifetime(time.Second * 10)
}

// GetMySQLConnection Get database connection
func GetMySQLConnection() *DB {
	return db
}
