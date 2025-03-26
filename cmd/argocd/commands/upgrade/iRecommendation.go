package upgrade

type IRecommendation interface {
	setName(name string)
	getName() string
	setServerVersion(version []int)
	getServerVersion() []int
	setUpgradeVersion(version []int)
	getUpgradeVersion() []int
}
