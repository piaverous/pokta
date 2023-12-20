package cmd

import (
	"github.com/piaverous/pokta/pokta"
	"github.com/spf13/cobra"
)

func New(app *pokta.App) *cobra.Command {
	return buildPoktaCommand(app)
}

func buildPoktaCommand(app *pokta.App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pokta",
		Short: "pokta helps you get info from your Jira projects",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return app.Config.Load(cmd.Flags())
		},
	}

	cmd.AddCommand(buildAuthCommand(app))
	cmd.AddCommand(buildConfigCommand(app))

	return cmd
}
