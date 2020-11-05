package conf

import (
	"database/sql"
	"fmt"
	"log"

	"gcrawler/model"

	_ "github.com/go-sql-driver/mysql"
)

/**
 * @Description: 初始化db
 * @Author: hunter
 * @Date: 2020-10-03 20:00
 * @LastEditTime: 2020-10-03 20:03:00
 * @LastEditors: hunter
 */
func InitDB(dsn string, maxIdleConn, maxOpenConn int) *sql.DB {

	db, err := sql.Open("mysql", dsn)
	fmt.Println(err)
	if err != nil {
		log.Fatalln(err)
		return db
	}

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

/**
 * @Description: 连接MYSQL数据库
 * @Author: hunter
 * @Date: 2020-10-27 12:31:12
 * @LastEditTime: 2020-10-27 12:32:45
 * @LastEditors: hunter
 */
func DbConnection() {
	//DB := InitDB("root:cjds1023456@tcp(172.21.34.72:3306)/live_program?charset=utf8mb4", 100, 600)
	DB := InitDB("root:c4f186b9e03812*!@tcp(127.0.0.1:3306)/live_program?charset=utf8mb4", 100, 600)
	model.Constructor(DB, nil, nil)
}
