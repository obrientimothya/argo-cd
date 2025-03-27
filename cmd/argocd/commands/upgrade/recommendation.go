package upgrade

type Recommendation interface {
	PrintRecommendation(r Upgrade)
}
