package cmd

import (
	"github.com/piaverous/pokta/pokta"
	"github.com/spf13/cobra"
)

func buildAuthCommand(app *pokta.App) *cobra.Command {
	auth := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate to Okta.",
	}

	auth.AddCommand(buildAuthPkjwtCommand(app))

	return auth
}
