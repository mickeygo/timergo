package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

var (
	DBDriver 	string = "sqlite3"
	DataSource 	string = "./db/timergo.db"
)

// Open 连接 db
func Open() (*sqlx.DB, error) {
	return sqlx.Open(DBDriver, DataSource)
}

// Connect 连接 db 并验证是否成功连接
func Connect() (*sqlx.DB, error) {
	return sqlx.Connect(DBDriver, DataSource)
}