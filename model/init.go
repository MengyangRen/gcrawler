package model

import (
	"database/sql"

	"github.com/go-redis/redis/v7"
)

var (
	pool *redis.Client
	db   *sql.DB
	api  []byte
)

func Constructor(sqldb *sql.DB, adminRedis *redis.Client, apiData []byte) {

	db = sqldb
	pool = adminRedis
	api = apiData

}
