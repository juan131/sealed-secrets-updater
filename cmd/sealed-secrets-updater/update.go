package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	"github.com/juan131/sealed-secrets-updater/pkg/config"
	"github.com/juan131/sealed-secrets-updater/pkg/updater"
)

var onlySecrets []string
var skipSecrets []string

// newCmdUpdate creates a command object for the "update" action.
func newCmdUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update your sealed secrets manifests.",
		Long:  "Track changes in your secrets manager and update your sealed secrets manifests.",
		Example: `
	sealed-secrets-updater update --config {config}   Update sealed secrets manifests based on {config} config file`,
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := config.New(configPath)
			if err != nil {
				return fmt.Errorf("unable to load config: %w", err)
			}

			if err := config.Validate(); err != nil {
				return fmt.Errorf("invalid config: %w", err)
			}

			if err := updater.UpdateSealedSecrets(
				context.Background(),
				config,
				updater.Filter{OnlySecrets: onlySecrets, SkipSecrets: skipSecrets},
			); err != nil {
				return fmt.Errorf("unable to update sealed secrets: %w", err)
			}

			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().StringSliceVar(&onlySecrets, "only-secrets", []string{}, "Only update provided list of secrets")
	cmd.Flags().StringSliceVar(&skipSecrets, "skip-secrets", []string{}, "List of secrets to skip updating")

	// Flags common to all sub commands
	cmd.PersistentFlags().StringVar(&configPath, "config", "", "Path to config file")
	if err := cmd.MarkPersistentFlagRequired("config"); err != nil {
		klog.Fatal(err)
	}

	return cmd
}
