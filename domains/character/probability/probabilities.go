package probability

import (
	"log"
	"math/rand"
	"time"
)

type CharacterProbabilities []CharacterProbability

func (ps CharacterProbabilities) CalculateOneHundred() CharacterProbabilities {
	//ランダムにシャッフル
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ps), func(i, j int) { ps[i], ps[j] = ps[j], ps[i] })

	results := make(CharacterProbabilities, 0, len(ps))
	var probability uint64
	for _, v := range ps {
		probability += v.Probability

		log.Println("index", probability, v.Probability)
		if probability > 100 {
			continue
		}
		results = append(results, v)
	}
	return results
}
