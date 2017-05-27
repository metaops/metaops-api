package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"time"
)

func init() {
	List = append(List, &gormigrate.Migration{
		ID: "201705270451_create_apps_table",
		Migrate: func(tx *gorm.DB) error {
			type App struct {
				ID   string `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
				Name string `gorm:"type:varchar(100)"`

				CreatedAt time.Time
				UpdatedAt time.Time
				DeletedAt *time.Time
			}
			err := tx.AutoMigrate(&App{}).Error

			return err
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("apps").Error
		},
	})
}
