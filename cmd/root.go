package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(
		dbCmd,
	)
}

var RootCmd = &cobra.Command{
	Use: "metaops",
	Run: runServerCmd,
}
