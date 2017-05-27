package cmd

import (
	"github.com/metaops/metaops-api/server"
	"github.com/spf13/cobra"
)

func runServerCmd(cmd *cobra.Command, args []string) {
	server.Init()
}
