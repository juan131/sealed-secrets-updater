package main

import (
	goflag "flag"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

var configPath string

func newCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "sealed-secrets-updater",
		Short: "sealed-secrets-updater updates your sealed secrets manifests.",
		Long:  "CLI for tracking changes in your secrets manager and updating your sealed secrets manifests",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				klog.Fatal(err)
			}
		},
	}

	// Subcommands
	cmd.AddCommand(newCmdUpdate())
	cmd.AddCommand(newCmdValidate())

	// Flags common to all sub commands
	cmd.PersistentFlags().StringVar(&configPath, "config", "", "Path to config file")
	if err := cmd.MarkPersistentFlagRequired("config"); err != nil {
		return nil, err
	}

	return cmd, nil
}

func main() {
	command, err := newCommand()
	if err != nil {
		klog.Fatal(err)
	}

	klog.InitFlags(nil)
	goflag.Parse()
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	if err := command.Execute(); err != nil {
		klog.Fatal(err)
	}
}
