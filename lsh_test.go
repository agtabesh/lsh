package main

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/agtabesh/lsh/hash_family"
	"github.com/agtabesh/lsh/similarity_measure"
	"github.com/agtabesh/lsh/store"
	"github.com/agtabesh/lsh/types"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	ctx := context.Background()
	config := getConfig()
	lsh, err := NewLSH(config)
	assert.Nil(t, err)

	hashFamily, err := hash_family.GetHashFamily(config.HashFamily, config.SignatureSize)
	assert.Nil(t, err)

	vectors := getSampleVectors()

	for i, vector := range vectors {
		vectorID := types.VectorID(fmt.Sprint(i))
		signature := hashFamily.MinHash(vector)
		err := lsh.Add(ctx, vectorID, vector)
		assert.Nil(t, err)

		storedSignature, err := lsh.store.GetSignatureByVectorID(ctx, vectorID)
		assert.Nil(t, err)

		assert.True(t, slices.Equal(signature, storedSignature))
	}
}

func TestQueryByVectorID(t *testing.T) {
	ctx := context.Background()
	config := getConfig()
	lsh, err := NewLSH(config)
	assert.Nil(t, err)

	vectors := getSampleVectors()

	for i, vector := range vectors {
		vectorID := types.VectorID(fmt.Sprint(i))
		err := lsh.Add(ctx, vectorID, vector)
		assert.Nil(t, err)
	}

	for i := range vectors {
		vectorID := types.VectorID(fmt.Sprint(i))
		vectorsID, err := lsh.QueryByVectorID(ctx, vectorID, 100)
		assert.Nil(t, err)
		assert.Contains(t, vectorsID, vectorID)
	}
}

func TestQueryByVector(t *testing.T) {
	ctx := context.Background()
	config := getConfig()
	lsh, err := NewLSH(config)
	assert.Nil(t, err)

	vectors := getSampleVectors()

	for i, vector := range vectors {
		vectorID := types.VectorID(fmt.Sprint(i))
		err := lsh.Add(ctx, vectorID, vector)
		assert.Nil(t, err)
	}

	for i, vector := range vectors {
		vectorID := types.VectorID(fmt.Sprint(i))
		vectorsID, err := lsh.QueryByVector(ctx, vector, 100)
		assert.Nil(t, err)
		assert.Contains(t, vectorsID, vectorID)
	}
}

func getConfig() LSHConfig {
	return LSHConfig{
		SignatureSize:     128,
		BandSize:          16,
		HashFamily:        hash_family.XXHash64,
		SimilarityMeasure: similarity_measure.HammingSimilarity,
		Store:             store.InMemoryStore,
	}
}

func getSampleVectors() []types.Vector {
	return []types.Vector{
		{"feat1": 1, "feat2": 1, "feat3": 1},
		{"feat1": 1, "feat4": 1, "feat5": 1},
		{"feat1": 1, "feat2": 1, "feat6": 1, "feat7": 1},
	}
}
