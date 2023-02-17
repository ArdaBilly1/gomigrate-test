package services

import (
	"gomigrate-test/migration/documents"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	if err := documents.InitMigration(db).Migrate(); err != nil {
		return err
	}

	return nil
}

func MigrateTo(db *gorm.DB, id string) error {
	if err := documents.InitMigration(db).MigrateTo(id); err != nil {
		return err
	}

	return nil
}
