package store

import (
	"context"

	"github.com/agtabesh/lsh/interfaces"
	"github.com/agtabesh/lsh/types"
)

var _ interfaces.Store = (*InMemoryStore)(nil)

const BUCKET_WINDOW_SIZE_START = 1

// InMemoryStore is a struct that implements the Store interface
type InMemoryStore struct {
	signaturesMap map[types.VectorID]types.Signature
	bucketsMap    map[types.VectorID]types.Buckets
}

// NewInMemoryStore creates a new instance of InMemoryStore
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		signaturesMap: make(map[types.VectorID]types.Signature),
		bucketsMap:    make(map[types.VectorID]types.Buckets),
	}
}

// GetSignatureByVectorID retrieves the signature for a given vector ID
func (s *InMemoryStore) GetSignatureByVectorID(ctx context.Context, vectorID types.VectorID) (types.Signature, error) {
	return s.signaturesMap[vectorID], nil
}

// GetVectorsIDInBucket retrieves vector IDs associated with given bucket IDs
func (s *InMemoryStore) GetVectorsIDInBucket(ctx context.Context, bucketID types.Bucket) (chan types.VectorID, error) {
	vectorIDMap := make(map[types.VectorID]bool)
	ch := make(chan types.VectorID)
	go func() {
		defer func() {
			vectorIDMap = make(map[types.VectorID]bool)
			close(ch)
		}()

		bucketWindowSize := BUCKET_WINDOW_SIZE_START
		i := 0
		for vID, bsID := range s.bucketsMap {
			for _, bID := range bsID {
				i += 1
				if i == bucketWindowSize {
					select {
					case <-ctx.Done():
						return
					default:
						bucketWindowSize *= 2
						i = 0
					}
				}
				if bID != bucketID {
					continue
				}
				if _, ok := vectorIDMap[vID]; ok {
					continue
				}
				ch <- vID
				vectorIDMap[vID] = true
			}
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
	s.bucketsMap[vectorID] = bucketsID
	return nil
}
