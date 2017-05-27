package cmd

import (
	"github.com/metaops/metaops-api/app"
	"github.com/metaops/metaops-api/config"
	"github.com/metaops/metaops-api/server"
	"github.com/spf13/cobra"
)

func runServerCmd(cmd *cobra.Command, args []string) {
	appConfig := config.Load()
	a := app.New(&appConfig)
	server := server.New(a, &appConfig.ServerConfig)

	server.Init()
}
