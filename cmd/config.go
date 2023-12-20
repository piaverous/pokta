package cmd

import (
	"github.com/piaverous/pokta/pokta"
	"github.com/spf13/cobra"
)

func buildConfigCommand(app *pokta.App) *cobra.Command {
	config := &cobra.Command{
		Use:   "config",
		Short: "Tools to configure pokta",
	}

	config.AddCommand(buildConfigShowCommand(app))
	return config
}
