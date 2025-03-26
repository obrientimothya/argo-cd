package upgrade

import (
	"fmt"
	"strconv"
	"strings"
)

func GetVersionTag(version string) string {
	return strings.Split(version, "+")[0]
}

func getMajorMinorVersion(version string) (int, int, error) {
	errMsg := "invalid upgrade version tag format"

	parts := strings.Split(strings.TrimPrefix(version, "v"), ".")
	if len(parts) < 2 {
		return 0, 0, fmt.Errorf("%s %s", errMsg, version)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("%s %s", errMsg, version)
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("%s %s", errMsg, version)
	}

	return major, minor, nil
}

func PrintRecommendations(serverVersion string, upgradeVersion string) error {
	fmt.Printf("Server Version   %s\n", serverVersion)
	fmt.Printf("Upgrade Version  %s\n", upgradeVersion)

	serverMajor, serverMinor, err := getMajorMinorVersion(serverVersion)
	if err != nil {
		return err
	}
	upgradeMajor, upgradeMinor, err := getMajorMinorVersion(upgradeVersion)
	if err != nil {
		return err
	}

	recommendationType := fmt.Sprintf("v%dv%d", serverMajor, upgradeMajor)
	recommendation, err := getRecommendation(recommendationType)
	if err != nil {
		fmt.Printf("No recommendations found for this upgrade.\n")
		fmt.Printf("Ensure the `argocd` command is up to date for latest recommendations.\n")
		return nil
	}
	recommendation.setServerVersion([]int{serverMajor, serverMinor})
	recommendation.setUpgradeVersion([]int{upgradeMajor, upgradeMinor})
	fmt.Printf("%s\n", recommendation.getName())

	return nil
}
