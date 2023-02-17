package documents

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migration202301182142() *gormigrate.Migration {
	m := new(gormigrate.Migration)
	m.ID = "202301182142"
	m.Migrate = func(tx *gorm.DB) error {
		type UserRole struct {
			Subrole string `json:"sub_role"`
		}

		return tx.AutoMigrate(&UserRole{})
	}

	m.Rollback = func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("user_roles")
	}

	return m
}
