package upgrade

type Recommendation struct {
	name           string
	serverVersion  []int
	upgradeVersion []int
}

func (r *Recommendation) setName(name string) {
	r.name = name
}

func (r *Recommendation) getName() string {
	return r.name
}

func (r *Recommendation) setServerVersion(serverVersion []int) {
	r.serverVersion = serverVersion
}

func (r *Recommendation) getServerVersion() []int {
	return r.serverVersion
}

func (r *Recommendation) setUpgradeVersion(upgradeVersion []int) {
	r.upgradeVersion = upgradeVersion
}

func (r *Recommendation) getUpgradeVersion() []int {
	return r.upgradeVersion
}
