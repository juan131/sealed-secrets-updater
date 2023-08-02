package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/juan131/sealed-secrets-updater/pkg/config"
)

// newCmdValidate creates a command object for the "validate" action.
func newCmdValidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate configuration.",
		Long:  "Validate sealed-secrets-updater configuration.",
		Example: `
	sealed-secrets-updater validate --config {config}   Validate {config} config file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := config.New(configPath)
			if err != nil {
				return fmt.Errorf("unable to load config: %w", err)
			}

			if err := config.Validate(); err != nil {
				return fmt.Errorf("invalid config: %w", err)
			}

			klog.Info("Config is valid! ✅")

			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	return cmd
}
