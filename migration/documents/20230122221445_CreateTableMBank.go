package documents

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// PRD Link :
// DDL link :
// Clickup Link :
func CreateTableMBank20230122221445() *gormigrate.Migration {
	m := new(gormigrate.Migration)

	m.ID = "20230122221445" // your id generated, please don't change
	m.Migrate = func(tx *gorm.DB) error {
		// replace this
		type Test struct {
			gorm.Model
			ID        string    `json:"id" gorm:"primaryKey"`
			Name      string    `json:"name"`
			CreatedAt time.Time `json:"created_at"`
		}

		return tx.AutoMigrate(&Test{})
	}

	m.Rollback = func(tx *gorm.DB) error {
		// replace this, what do you want
		return tx.Migrator().DropTable("test")
	}

	return m
}
