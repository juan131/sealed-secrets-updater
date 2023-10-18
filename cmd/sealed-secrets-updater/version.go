package main

import (
	"github.com/spf13/cobra"
)

// newCmdVersion creates a command object for the "version" action.
func newCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print sealed-secrets-updater version",
		Run: func(cmd *cobra.Command, args []string) {
			root := cmd.Root()
			root.SetArgs([]string{"--version"})
			root.Execute()
		},
	}

	return cmd
}
