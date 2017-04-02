package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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
	con, err := sql.Open("mysql", "root:123456@/shuttlebus")
	if err != nil {
		panic(err.Error())
	}
	err = con.Ping()
	if err != nil {
		panic(err.Error())
	}
	return con
}
