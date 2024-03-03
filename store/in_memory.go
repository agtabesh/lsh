package store

import (
	"context"

	"github.com/agtabesh/lsh/interfaces"
	"github.com/agtabesh/lsh/types"
)

var _ interfaces.Store = (*inMemoryStore)(nil)

// inMemoryStore is a struct that implements the Store interface
type inMemoryStore struct {
	signaturesMap map[types.VectorID]types.Signature
	bucketsMap    map[types.VectorID]types.Buckets
}

// NewInMemoryStore creates a new instance of InMemoryStore
func newInMemoryStore() *inMemoryStore {
	return &inMemoryStore{
		signaturesMap: make(map[types.VectorID]types.Signature),
		bucketsMap:    make(map[types.VectorID]types.Buckets),
	}
}

// GetSignatureByVectorID retrieves the signature for a given vector ID
func (s *inMemoryStore) GetSignatureByVectorID(ctx context.Context, vectorID types.VectorID) (types.Signature, error) {
	return s.signaturesMap[vectorID], nil
}

// GetVectorsIDInBuckets retrieves vector IDs associated with given bucket IDs
func (s *inMemoryStore) GetVectorsIDInBuckets(ctx context.Context, bucketsID types.Buckets) ([]types.VectorID, error) {
	vectorsMap := make(map[types.VectorID]bool)
	for _, bucketID := range bucketsID {
		for vectorID, bucketID2 := range s.bucketsMap {
			for _, b1 := range bucketID2 {
				if b1 != bucketID {
					continue
				}
				if vectorsMap[vectorID] {
					continue
				}
				vectorsMap[vectorID] = true
			}
		}
	}
	vectorsID := make([]types.VectorID, 0)
	for vectorID := range vectorsMap {
		vectorsID = append(vectorsID, vectorID)
	}

	return vectorsID, nil
}

// UpdateSignatureByVectorID updates the signature for a given vector ID
func (s *inMemoryStore) UpdateSignatureByVectorID(ctx context.Context, vectorID types.VectorID, signature types.Signature) error {
	s.signaturesMap[vectorID] = signature
	return nil
}

// UpdateBucketsByVectorID updates the buckets associated with a given vector ID
func (s *inMemoryStore) UpdateBucketsByVectorID(ctx context.Context, vectorID types.VectorID, bucketsID types.Buckets) error {
	s.bucketsMap[vectorID] = bucketsID
	return nil
}
