package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"time"
)

func init() {
	List = append(List, &gormigrate.Migration{
		ID: "201705280253_create_node_deployments_table",
		Migrate: func(tx *gorm.DB) error {
			type NodeDeployment struct {
				ID           string `gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
				DeploymentID string `gorm:"type:uuid"`
				NodeID       string `gorm:"type:uuid"`
				Status       string `gorm:"type:varchar(100)"`

				CreatedAt time.Time
				UpdatedAt time.Time
				DeletedAt *time.Time
			}

			if err := tx.AutoMigrate(&NodeDeployment{}).Error; err != nil {
				return err
			}

			if err := tx.Model(NodeDeployment{}).AddForeignKey("deployment_id", "deployments (id)", "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}

			err := tx.Model(NodeDeployment{}).AddForeignKey("node_id", "nodes (id)", "RESTRICT", "RESTRICT").Error

			return err

		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("node_deployments").Error
		},
	})
}
