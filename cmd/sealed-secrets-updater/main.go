package main

import (
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var (
	configPath string
	Version    string
)

func newCommand() (*cobra.Command, error) {
	if Version == "" {
		Version = "dev"
	}

	cmd := &cobra.Command{
		Use:     "sealed-secrets-updater",
		Short:   "sealed-secrets-updater updates your sealed secrets manifests.",
		Long:    "CLI for tracking changes in your secrets manager and updating your sealed secrets manifests",
		Version: Version,
	}

	// Set version template
	cmd.SetVersionTemplate(`{{ printf "%s\n" .Version }}`)

	// Subcommands
	cmd.AddCommand(newCmdUpdate())
	cmd.AddCommand(newCmdValidate())
	cmd.AddCommand(newCmdVersion())

	return cmd, nil
}

func main() {
	command, err := newCommand()
	if err != nil {
		klog.Fatal(err)
	}

	if err := command.Execute(); err != nil {
		klog.Fatal(err)
	}
}
