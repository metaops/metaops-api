package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"time"
)

func init() {
	List = append(List, &gormigrate.Migration{
		ID: "201705270647_create_nodes_table",
		Migrate: func(tx *gorm.DB) error {
			type Node struct {
				ID    string `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
				AppID string `gorm:"type:uuid"`

				CreatedAt time.Time
				UpdatedAt time.Time
				DeletedAt *time.Time
			}

			if err := tx.AutoMigrate(&Node{}).Error; err != nil {
				return err
			}

			err := tx.Model(Node{}).AddForeignKey("app_id", "apps (id)", "RESTRICT", "RESTRICT").Error

			return err

		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("nodes").Error
		},
	})
}
