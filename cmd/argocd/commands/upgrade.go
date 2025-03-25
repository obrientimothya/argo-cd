package commands

import (
	argocdclient "github.com/argoproj/argo-cd/v3/pkg/apiclient"
	"github.com/spf13/cobra"
	"os"
)

// NewUpgradeCmd returns a new `upgrade` command to be used as a sub-command to root
func NewUpgradeCmd(clientOpts *argocdclient.ClientOptions) *cobra.Command {
	var (
		tag string
	)

	upgradeCmd := cobra.Command{
		Use:   "upgrade",
		Short: "Prints configuration changes recommended to prepare your install for an upgrade",
		Example: `  # Print recommendations for the latest 'stable' release
  argocd upgrade

  # Print recommendations for a specific release. See: https://github.com/argoproj/argo-cd/releases
  argocd upgrade --tag v3.0.0
`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
			os.Exit(1)
		},
	}
	upgradeCmd.Flags().StringVarP(&tag, "tag", "t", "stable", "Release tag to check for upgrade recommendations. See: https://github.com/argoproj/argo-cd/releases")
	return &upgradeCmd
}
