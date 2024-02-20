package interfaces

import "github.com/agtabesh/lsh/internal/types"

type SimilarityMeasure interface {
	SetPrecision(p int8)
	Calculate(x, y types.Quantitative) float64
}
