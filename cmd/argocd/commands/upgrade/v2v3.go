package upgrade

import "fmt"

type V2V3Recommendation struct{}

func (v *V2V3Recommendation) PrintRecommendation(u Upgrade) {
	fmt.Printf("SERVER: %s\n", u.not(SERVER, 2.14))
}
