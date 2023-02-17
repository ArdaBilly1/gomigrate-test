package migration

import (
	"gomigrate-test/migration/common"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func OpenConnectionMysql() {
	common.GenerateInfo(common.MsgOpenDatabaseMysql)
	dsn := "root:root@tcp(127.0.0.1:8889)/testgorm?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		common.GenerateError(err.Error())
		os.Exit(0)
	}

	db = conn
}

func OpenConnectionPostgres() {

}
