package lsh

import (
	"context"

	"github.com/agtabesh/lsh/interfaces"
	"github.com/agtabesh/lsh/types"
)

// LSHConfig holds configuration parameters for LSH.
type LSHConfig struct {
	SignatureSize int // SignatureSize is the size of the signature.
}

// LSH represents the Locality Sensitive Hashing service.
type LSH struct {
	config            LSHConfig                    // Configuration for LSH.
	hashFamily        interfaces.HashFamily        // HashFamily is the hash family for LSH.
	similarityMeasure interfaces.SimilarityMeasure // SimilarityMeasure is the measure used for similarity computation.
	store             interfaces.Store             // Store is the data store for LSH.
}

// NewLSH creates a new instance of LSH.
func NewLSH(
	config LSHConfig,
	hashFamily interfaces.HashFamily,
	similarityMeasure interfaces.SimilarityMeasure,
	store interfaces.Store,
) (*LSH, error) {
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
	buckets := signature.Buckets()

	// Update the signature and buckets in the store.
	s.store.UpdateBucketsByVectorID(ctx, vectorID, buckets)
	s.store.UpdateSignatureByVectorID(ctx, vectorID, signature)

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
	bucketCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Calculate buckets for the provided signature.
	buckets := signature.Buckets()
	similarities := make(types.Similarities)
	for _, bucket := range buckets {
		// Get candidate vectors from the buckets.
		candidateVectorsIDChan, err := s.store.GetVectorsIDInBucket(bucketCtx, bucket)
		if err != nil {
			return []types.VectorID{}, err
		}

		// Calculate similarities between the query signature and candidate signatures.
		for candidateID := range candidateVectorsIDChan {
			if _, ok := similarities[candidateID]; ok {
				continue
			}
			candidateSignature, err := s.store.GetSignatureByVectorID(ctx, candidateID)
			if err != nil {
				return []types.VectorID{}, err
			}

			similarities[candidateID] = s.similarityMeasure.Measure(signature, candidateSignature)
		}
		if len(similarities) >= count {
			cancel()
			break
		}
	}

	// Return top similar vectors.
	return similarities.Top(count), nil
}
