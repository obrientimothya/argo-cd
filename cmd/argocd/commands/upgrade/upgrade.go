package upgrade

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	SERVER  int = 0
	UPGRADE int = 1
)

type Upgrade struct {
	ServerVersion  []int
	UpgradeVersion []int
}

func (u *Upgrade) getServerVersion() []int {
	return u.ServerVersion
}

func (u *Upgrade) getUpgradeVersion() []int {
	return u.UpgradeVersion
}

func (u *Upgrade) getVersionFloat(t int) float64 {
	var version []int
	if t == 0 {
		version = u.ServerVersion
	} else {
		version = u.UpgradeVersion
	}
	return versionToFloat(version)
}

func (u *Upgrade) lte(t int, version float64) bool {
	if u.getVersionFloat(t) <= version {
		return true
	}
	return false
}

func (u *Upgrade) gte(t int, version float64) bool {
	if u.getVersionFloat(t) >= version {
		return true
	}
	return false
}

func (u *Upgrade) lt(t int, version float64) bool {
	if u.getVersionFloat(t) < version {
		return true
	}
	return false
}

func (u *Upgrade) gt(t int, version float64) bool {
	if u.getVersionFloat(t) > version {
		return true
	}
	return false
}

func (u *Upgrade) eq(t int, version float64) bool {
	if u.getVersionFloat(t) == version {
		return true
	}
	return false
}

func (u *Upgrade) not(t int, version float64) bool {
	if u.getVersionFloat(t) != version {
		return true
	}
	return false
}

func Run(serverVersion string, upgradeVersion string) error {
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
	recommendation, err := GetRecommendation(recommendationType)
	if err != nil {
		fmt.Printf("No recommendations found for this upgrade.\n")
		fmt.Printf("Ensure the `argocd` command is up to date for latest recommendations.\n")
		return nil
	}

	recommendation.PrintRecommendation(Upgrade{
		ServerVersion:  []int{serverMajor, serverMinor},
		UpgradeVersion: []int{upgradeMajor, upgradeMinor},
	})

	return nil
}

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

func versionToFloat(version []int) float64 {
	digitCount := int(math.Log10(float64(version[1]))) + 1
	fraction := float64(version[1]) / math.Pow(10, float64(digitCount))
	return float64(version[0]) + fraction
}
