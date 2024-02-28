package interfaces

import "github.com/agtabesh/lsh/internal/types"

type SimilarityMeasure interface {
	Measure(x, y types.Signature) float64
}
