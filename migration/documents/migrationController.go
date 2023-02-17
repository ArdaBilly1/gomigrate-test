package documents

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		Migration202301170830(),
		Migration202301182142(),
	})
}
