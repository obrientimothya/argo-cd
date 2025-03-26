package commands

import (
	"fmt"
	"github.com/argoproj/argo-cd/v3/cmd/argocd/commands/upgrade"
	"github.com/argoproj/argo-cd/v3/common"
	argocdclient "github.com/argoproj/argo-cd/v3/pkg/apiclient"
	"github.com/argoproj/argo-cd/v3/util/errors"
	"github.com/spf13/cobra"
)

// NewUpgradeCmd returns a new `upgrade` command to be used as a sub-command to root
func NewUpgradeCmd(clientOpts *argocdclient.ClientOptions) *cobra.Command {
	var tag string
	clientVersion := upgrade.GetVersionTag(common.GetVersion().String())

	upgradeCmd := cobra.Command{
		Use:   "upgrade",
		Short: "Prints configuration changes recommended to prepare your Server for an upgrade",
		Example: fmt.Sprintf(`  # Print recommendations for upgrading server to the current client version (%s)
  argocd upgrade

  # Print recommendations for a specific release tag. See: https://github.com/argoproj/argo-cd/releases
  argocd upgrade --tag %s
`, clientVersion, clientVersion),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			serverVersion := upgrade.GetVersionTag(getServerVersion(ctx, clientOpts, cmd).Version)

			if tag != "" {
				clientVersion = tag
			}

			err := upgrade.PrintRecommendations(serverVersion, clientVersion)
			if err != nil {
				errors.CheckError(err)
			}
		},
	}
	upgradeCmd.Flags().StringVarP(&tag, "tag", "t", clientVersion, "Release tag to check for upgrade recommendations. See: https://github.com/argoproj/argo-cd/releases")
	return &upgradeCmd
}
