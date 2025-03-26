package upgrade

import (
	"fmt"
)

func getRecommendation(RecommendationType string) (IRecommendation, error) {
	if RecommendationType == "v2v3" {
		return newV2V3(), nil
	}
	return nil, fmt.Errorf("RecommendationNotFound")
}
