package cmd

import (
	"fmt"
	"os"

	"github.com/piaverous/pokta/pokta"
	"github.com/spf13/cobra"
)

func buildAuthPkjwtCommand(app *pokta.App) *cobra.Command {
	var SignOnly bool
	cmd := &cobra.Command{
		Use:   "pkjwt [flags]",
		Short: "Authenticate using PKJWT",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := app.AuthenticatePKJWT(SignOnly)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			fmt.Print(response)
			return nil
		},
	}
	cmd.PersistentFlags().BoolVarP(&SignOnly, "sign-only", "s", false, "If true, only sign JWT with private key locally and do not fetch an access_token from Okta.")

	return cmd
}
