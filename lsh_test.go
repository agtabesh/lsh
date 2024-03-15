package lsh

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/agtabesh/lsh/interfaces"
	"github.com/agtabesh/lsh/types"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	ctx := context.Background()
	lsh, hashFamily, _, _, err := getInstance()
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
	lsh, _, _, _, err := getInstance()
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
	lsh, _, _, _, err := getInstance()
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
		SignatureSize: 128,
	}
}

func getInstance() (
	*LSH,
	interfaces.HashFamily,
	interfaces.SimilarityMeasure,
	interfaces.Store,
	error,
) {
	config := getConfig()
	hashFamily := NewXXHASH64HashFamily(config.SignatureSize)
	similarityMeasure := NewHammingSimilarity()
	store := NewInMemoryStore()

	lsh, err := NewLSH(config, hashFamily, similarityMeasure, store)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return lsh, hashFamily, similarityMeasure, store, nil
}

func getSampleVectors() []types.Vector {
	return []types.Vector{
		{"feat1": 1, "feat2": 1, "feat3": 1},
		{"feat1": 1, "feat4": 1, "feat5": 1},
		{"feat1": 1, "feat2": 1, "feat6": 1, "feat7": 1},
	}
}
