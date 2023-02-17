package common

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

var (
	MsgMigrationCommandNotFound = "Migration command not found!"
	MsgRollbackCommandNotFound  = "Rollback command not found!"

	MsgMigrationError   = "Error when migrate"
	MsgMigrationSuccess = "Migration Success"

	MsgRollbackError   = "Error when rollback"
	MsgRollbackSuccess = "Rollback Success"

	MsgOpenDatabaseMysql = "Open connection Database Mysql...."

	MsgErrGenerateFileMigration = "Unable to parse data into template: %v\n"
	MsgGeneratorFileSuccess     = "Migration file created, please check at %s"
)

func ShowBanner() {
	banner := figure.NewFigure("gormigrate", "", true)
	banner.Print()
	fmt.Println()
	fmt.Println()
}

func GenerateInfo(msg string) {
	fmt.Println("INFO:", msg)
}

func GenerateWarning(msg string) {
	fmt.Println("WARNING:", msg)
}

func GenerateError(msg string) {
	fmt.Println("ERROR:", msg)
}
