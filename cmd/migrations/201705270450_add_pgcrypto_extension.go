package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	List = append(List, &gormigrate.Migration{
		ID: "201705270450_add_pgcrypto_extension",
		Migrate: func(tx *gorm.DB) error {
			return tx.Exec("CREATE EXTENSION pgcrypto;").Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Exec("DROP EXTENSION IF EXISTS pgcrypto;").Error
		},
	})
}
