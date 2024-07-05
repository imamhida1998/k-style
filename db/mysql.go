package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	MySQL *sql.DB
)

func InitDB() (err error) {
	url := "admin:wRYS452im2yylSWqVlX8@tcp(developer.cfw42e08oubz.us-east-1.rds.amazonaws.com:3306)/test?parseTime=true"

	MySQL, err = sql.Open("mysql", url)
	if err != nil {
		fmt.Print("err: ", err.Error())
		return err
	}

	err = MySQL.Ping()
	if err != nil {
		fmt.Print("err: ", err.Error())
		return err
	}

	return nil
}
