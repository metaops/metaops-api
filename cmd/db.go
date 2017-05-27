package cmd

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/metaops/metaops-api/cmd/migrations"
	"github.com/metaops/metaops-api/config"
	"github.com/spf13/cobra"
	"gopkg.in/gormigrate.v1"
)

func init() {
	dbCmd.AddCommand(
		dbMigrateCmd,
		dbRollbackCmd,
	)
}

var dbCmd = &cobra.Command{
	Use: "db",
}

var dbMigrateCmd = &cobra.Command{
	Use: "migrate",
	Run: dbMigrateCmdFunc,
}

var dbRollbackCmd = &cobra.Command{
	Use: "rollback",
	Run: dbRollbackCmdFunc,
}

func dbMigrateCmdFunc(cmd *cobra.Command, args []string) {
	m := initMigrations(cmd)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")
}

func dbRollbackCmdFunc(cmd *cobra.Command, args []string) {
	m := initMigrations(cmd)
	if err := m.RollbackLast(); err != nil {
		log.Fatalf("Could not rollback: %v", err)
	}

	log.Printf("Rollback did run successfully")
}

func initMigrations(cmd *cobra.Command) *gormigrate.Gormigrate {

	dbConfig := config.Load().DBConfig

	db, err := gorm.Open(dbConfig.Driver, dbConfig.Connection)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	return gormigrate.New(db, gormigrate.DefaultOptions, migrations.List)
}
