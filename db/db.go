package db

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

func Query(query string, args ...interface{}) *sql.Rows {
	con := getConnection()
	defer con.Close()
	r, err := con.Query(query, args...)
	if err != nil {
		panic(err.Error())
	}
	return r
}

func ExecuteWithError(sql string, args ...interface{}) (sql.Result, error) {
	con := getConnection()
	r, e := con.Exec(sql, args...)
	return r, e
}

func Execute(sql string, args ...interface{}) sql.Result {
	con := getConnection()
	r, e := con.Exec(sql, args...)
	if e != nil {
		panic(e.Error())
	}
	return r
}

func getConnection() *sql.DB {
	con, err := sql.Open("mysql", "root:123456@/transportation")
	if err != nil {
		panic(err.Error())
	}
	err = con.Ping()
	if err != nil {
		panic(err.Error())
	}
	return con
}

func GenerateID() string {
	return uuid.NewV4().String()
}

func deleteRow(tableName, id string) {
	Execute(fmt.Sprintf("DELETE FROM %s WHERE ID = ?", tableName), id)
}
