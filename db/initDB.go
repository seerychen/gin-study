/*
 * @Author: cyr
 * @Version:
 * @Date: 2023-01-16 10:14:53
 * @Description:
 */
package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	var err error
	DbConnection, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/gin?parseTime=true")
	if err != nil {
		log.Panicln("err:", err.Error())
		panic(err.Error())
	} else {
		DbConnection.SetMaxOpenConns(10)
		DbConnection.SetMaxIdleConns(10)
		return DbConnection
	}
}
