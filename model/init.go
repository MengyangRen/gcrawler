package model

import (
	"database/sql"

	"github.com/go-redis/redis/v7"
)

var (
	pool  *redis.Client
	db    *sql.DB
	api []byte
)

/**
 * @Description: 初始化
 * @Author: hunter
 * @Date: 2020-10-07 11:34:20
 * @LastEditTime: 2020-10-07 11:36:16
 * @LastEditors: hunter
 */
func Constructor(sqldb *sql.DB, adminRedis *redis.Client, apiData []byte) {

	db = sqldb
	pool = adminRedis
	api = apiData

}

