package upgrade

import (
	"fmt"
)

func GetRecommendation(recommendationType string) (Recommendation, error) {
	switch recommendationType {
	case "v2v3":
		return new(V2V3Recommendation), nil
	default:
		return nil, fmt.Errorf("RecommendationNotFound")
	}
}
