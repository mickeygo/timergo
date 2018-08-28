package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	DBDriver string = "mysql"
)

// 获取数据库引擎
func GetEngine() (*xorm.Engine, error) {
	dataSource := ""
	return xorm.NewEngine(DBDriver, dataSource)
}