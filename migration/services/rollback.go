package services

import (
	"gomigrate-test/migration/documents"

	"gorm.io/gorm"
)

func RollbackTo(db *gorm.DB, id string) error {
	if err := documents.InitMigration(db).RollbackTo(id); err != nil {
		return err
	}

	return nil
}

func RollbackLast(db *gorm.DB) error {
	if err := documents.InitMigration(db).RollbackLast(); err != nil {
		return err
	}

	return nil
}
