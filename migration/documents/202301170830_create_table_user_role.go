package documents

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// PRD Link :
// DDL link :
// clickup Link :
func Migration202301170830() *gormigrate.Migration {
	m := new(gormigrate.Migration)

	m.ID = "202301170830"
	m.Migrate = func(tx *gorm.DB) error {
		type UserRole struct {
			gorm.Model
			ID        string    `json:"id" gorm:"primaryKey"`
			Role      string    `json:"role"`
			CreatedAt time.Time `json:"created_at"`
		}

		return tx.AutoMigrate(&UserRole{})
	}

	m.Rollback = func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("user_role")
	}

	return m
}
