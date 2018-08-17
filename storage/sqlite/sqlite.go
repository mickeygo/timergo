package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/xorm"
)

var (
	DBDriver string = "sqlite3"
)

// 获取数据库引擎
func GetEngine() (*xorm.Engine, error) {
	dataSource := ""
	return xorm.NewEngine(DBDriver, dataSource)
}