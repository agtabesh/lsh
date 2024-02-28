package similarity_measure

import (
	"math"

	"github.com/agtabesh/lsh/internal/interfaces"
	"github.com/agtabesh/lsh/internal/types"
)

var _ interfaces.SimilarityMeasure = (*HammingSimilarity)(nil)

const percision = 8

type HammingSimilarity struct{}

func NewHammingSimilarity() *HammingSimilarity {
	return &HammingSimilarity{}
}

// Calculate computes the hamming similarity between two signature x and y.
// It returns the similarity as an integer value between 0 and 1.
// The function assumes that x and y contain numeric values.
func (sm *HammingSimilarity) Measure(x, y types.Signature) float64 {
	if len(x) != len(y) {
		return 0
	}

	total := len(x)
	same := 0
	for i := 0; i < total; i += 1 {
		if x[i] == y[i] {
			same += 1
		}
	}

	similarity := float64(same) / float64(total)
	return math.Round(math.Pow(10, float64(percision))*similarity) / math.Pow(10, float64(percision))
}
