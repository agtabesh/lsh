package lsh

import (
	"github.com/agtabesh/lsh/hash_family"
	"github.com/agtabesh/lsh/interfaces"
	"github.com/agtabesh/lsh/similarity_measure"
	"github.com/agtabesh/lsh/store"
)

func NewXXHASH64HashFamily(count int) interfaces.HashFamily {
	return hash_family.NewXXHASH64HashFamily(count)
}

func NewHammingSimilarity() interfaces.SimilarityMeasure {
	return similarity_measure.NewHammingSimilarity()
}

func NewInMemoryStore() interfaces.Store {
	return store.NewInMemoryStore()
}
