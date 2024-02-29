package similarity_measure

import (
	"fmt"

	"github.com/agtabesh/lsh/internal/interfaces"
)

type SimilarityMeasure string

const HammingSimilarity SimilarityMeasure = "HAMMING_SIMILARITY_MEASURE"

func GetSimilarityMeasure(name SimilarityMeasure) (interfaces.SimilarityMeasure, error) {
	if name == HammingSimilarity {
		return newHammingSimilarity(), nil
	} else {
		return nil, fmt.Errorf("invalid similarity measure: %s", name)
	}
}
