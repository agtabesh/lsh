package store

import (
	"context"

	"github.com/agtabesh/lsh/interfaces"
	"github.com/agtabesh/lsh/types"
)

var _ interfaces.Store = (*InMemoryStore)(nil)

// InMemoryStore is a struct that implements the Store interface
type InMemoryStore struct {
	signaturesMap map[types.VectorID]types.Signature
	bucketsMap    map[types.Bucket]types.VectorsID
}

// NewInMemoryStore creates a new instance of InMemoryStore
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		signaturesMap: make(map[types.VectorID]types.Signature),
		bucketsMap:    make(map[types.Bucket]types.VectorsID),
	}
}

// GetSignatureByVectorID retrieves the signature for a given vector ID
func (s *InMemoryStore) GetSignatureByVectorID(ctx context.Context, vectorID types.VectorID) (types.Signature, error) {
	return s.signaturesMap[vectorID], nil
}

// GetVectorsIDInBucket retrieves vector IDs associated with given bucket IDs
func (s *InMemoryStore) GetVectorsIDInBucket(ctx context.Context, bucketID types.Bucket) (chan types.VectorID, error) {
	ch := make(chan types.VectorID)
	go func() {
		defer close(ch)
		for _, vID := range s.bucketsMap[bucketID] {
			ch <- vID
		}
	}()
	return ch, nil
}

// UpdateSignatureByVectorID updates the signature for a given vector ID
func (s *InMemoryStore) UpdateSignatureByVectorID(ctx context.Context, vectorID types.VectorID, signature types.Signature) error {
	s.signaturesMap[vectorID] = signature
	return nil
}

// UpdateBucketsByVectorID updates the buckets associated with a given vector ID
func (s *InMemoryStore) UpdateBucketsByVectorID(ctx context.Context, vectorID types.VectorID, bucketsID types.Buckets) error {
	bucketsToDelete := make(types.Buckets, 0)
	if _, ok := s.signaturesMap[vectorID]; ok {
		bucketsToDelete = s.signaturesMap[vectorID].Buckets()
	}

	for _, bID := range bucketsID.Diff(bucketsToDelete) {
		s.bucketsMap[bID] = append(s.bucketsMap[bID], vectorID)
	}

	for _, bID := range bucketsToDelete.Diff(bucketsID) {
		if m, ok := s.bucketsMap[bID]; ok {
			m.Remove(vectorID)
			s.bucketsMap[bID] = m
		}
	}
	return nil
}
