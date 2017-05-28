package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"time"
)

func init() {
	List = append(List, &gormigrate.Migration{
		ID: "201705280253_create_deployments_table",
		Migrate: func(tx *gorm.DB) error {
			type Deployment struct {
				ID     string `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
				AppID  string `gorm:"type:uuid"`
				Status string `gorm:"type:varchar(100)"`

				CreatedAt time.Time
				UpdatedAt time.Time
				DeletedAt *time.Time
			}

			if err := tx.AutoMigrate(&Deployment{}).Error; err != nil {
				return err
			}

			err := tx.Model(Deployment{}).AddForeignKey("app_id", "apps (id)", "RESTRICT", "RESTRICT").Error

			return err

		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("deployments").Error
		},
	})
}
