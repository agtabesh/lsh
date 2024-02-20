package similarity_measure

import (
	"math"

	"github.com/agtabesh/lsh/internal/interfaces"
	"github.com/agtabesh/lsh/internal/types"
)

var _ interfaces.SimilarityMeasure = (*CosineSimilarityMeasure)(nil)

type CosineSimilarityMeasure struct {
	precision int8
}

func NewCosineSimilarityMeasure(precision int8) *CosineSimilarityMeasure {
	return &CosineSimilarityMeasure{
		precision: precision,
	}
}

// SetPrecision sets the precision (number of decimal places) for rounding the similarity value.
func (sm *CosineSimilarityMeasure) SetPrecision(p int8) {
	sm.precision = p
}

// Calculate computes the cosine similarity between two quantitative data sets x and y.
// It returns the similarity as an integer value between -1 and 1.
// The function assumes that x and y contain numeric values, but it does not enforce that they are of the same length.
func (sm *CosineSimilarityMeasure) Calculate(x, y types.Quantitative) float64 {
	numerator, denominatorX, denominatorY := 0.0, 0.0, 0.0

	// Calculate the numerator of the cosine similarity formula
	for k, vx := range x {
		if vy, ok := y[k]; ok {
			numerator += vx * vy
		}
	}

	// Calculate the denominator of the cosine similarity formula
	for _, vx := range x {
		denominatorX += math.Pow(vx, 2)
	}

	for _, vy := range y {
		denominatorY += math.Pow(vy, 2)
	}
	denominator := math.Sqrt(denominatorX) * math.Sqrt(denominatorY)
	if denominator == 0 {
		return 0
	}

	// Calculate the cosine similarity
	similarity := numerator / denominator
	return math.Round(math.Pow(10, float64(sm.precision))*similarity) / math.Pow(10, float64(sm.precision))
}
