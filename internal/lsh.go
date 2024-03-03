package internal

import (
	"context"

	"github.com/agtabesh/lsh/internal/hash_family"
	"github.com/agtabesh/lsh/internal/interfaces"
	"github.com/agtabesh/lsh/internal/similarity_measure"
	"github.com/agtabesh/lsh/internal/store"
	"github.com/agtabesh/lsh/internal/types"
)

// LSHConfig holds configuration parameters for LSH.
type LSHConfig struct {
	SignatureSize     int                                  // SignatureSize is the size of the signature.
	BandSize          int                                  // BandSize is the size of the band for hashing.
	HashFamily        hash_family.HashFamily               // HashFamily is the hash family for LSH.
	SimilarityMeasure similarity_measure.SimilarityMeasure // SimilarityMeasure is the measure used for similarity computation.
	Store             store.Store                          // Store is the data store for LSH.
}

// LSH represents the Locality Sensitive Hashing service.
type LSH struct {
	config            LSHConfig                    // Configuration for LSH.
	hashFamily        interfaces.HashFamily        // HashFamily is the hash family for LSH.
	similarityMeasure interfaces.SimilarityMeasure // SimilarityMeasure is the measure used for similarity computation.
	store             interfaces.Store             // Store is the data store for LSH.
}

// NewLSH creates a new instance of LSH.
func NewLSH(config LSHConfig) (*LSH, error) {
	hashFamily, err := hash_family.GetHashFamily(config.HashFamily, config.SignatureSize)
	if err != nil {
		return nil, err
	}

	similarityMeasure, err := similarity_measure.GetSimilarityMeasure(config.SimilarityMeasure)
	if err != nil {
		return nil, err
	}

	store, err := store.GetStore(config.Store)
	if err != nil {
		return nil, err
	}
	return &LSH{
		config:            config,
		hashFamily:        hashFamily,
		similarityMeasure: similarityMeasure,
		store:             store,
	}, nil
}

// Add adds a vector to the LSH service.
func (s *LSH) Add(ctx context.Context, vectorID types.VectorID, vector types.Vector) error {
	// Get the signature for the vector ID.
	signature, err := s.store.GetSignatureByVectorID(ctx, vectorID)
	if err != nil {
		return err
	}

	// Update the signature with the new vector.
	for k := range vector {
		newSignature := s.hashFamily.Hash(k.String())
		signature = signature.FindMin(newSignature)
	}

	// Compute buckets for the updated signature.
	buckets := signature.Buckets(s.config.BandSize)

	// Update the signature and buckets in the store.
	s.store.UpdateSignatureByVectorID(ctx, vectorID, signature)
	s.store.UpdateBucketsByVectorID(ctx, vectorID, buckets)

	return nil
}

// QueryByVectorID performs a query using a vectorID.
func (s *LSH) QueryByVectorID(ctx context.Context, vectorID types.VectorID, count int) ([]types.VectorID, error) {
	// Get the signature for the vectorID.
	signature, err := s.store.GetSignatureByVectorID(ctx, vectorID)
	if err != nil {
		return []types.VectorID{}, err
	}
	return s.query(ctx, signature, count)
}

// QueryByVector performs a query using a vector.
func (s *LSH) QueryByVector(ctx context.Context, vector types.Vector, count int) ([]types.VectorID, error) {
	// Calculate the signature for the vector.
	var signature types.Signature
	for k := range vector {
		newSignature := s.hashFamily.Hash(k.String())
		signature = signature.FindMin(newSignature)
	}
	return s.query(ctx, signature, count)
}

// query performs the actual query using the provided signature.
func (s *LSH) query(ctx context.Context, signature types.Signature, count int) ([]types.VectorID, error) {
	// Compute buckets for the provided signature.
	buckets := signature.Buckets(s.config.BandSize)

	// Get candidate vectors from the buckets.
	candidateVectorsID, err := s.store.GetVectorsIDInBuckets(ctx, buckets)
	if err != nil {
		return []types.VectorID{}, err
	}

	// Compute similarities between the query signature and candidate signatures.
	similarities := make(types.Similarities)
	for _, candidateID := range candidateVectorsID {
		candidateSignature, err := s.store.GetSignatureByVectorID(ctx, candidateID)
		if err != nil {
			return []types.VectorID{}, err
		}

		similarities[candidateID] = s.similarityMeasure.Measure(signature, candidateSignature)
	}

	// Return top similar vectors.
	return similarities.Top(count), nil
}
