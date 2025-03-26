package upgrade

type V2V3 struct {
	Recommendation
}

func newV2V3() IRecommendation {
	return &V2V3{
		Recommendation: Recommendation{
			name: "Upgrade Recommendations v2 to v3",
		},
	}
}
