package interfaces

import (
	"context"

	"github.com/agtabesh/lsh/types"
)

type Store interface {
	GetSignatureByVectorID(ctx context.Context, vectorID types.VectorID) (types.Signature, error)
	GetVectorsIDInBuckets(ctx context.Context, bucketsID types.Buckets) ([]types.VectorID, error)
	UpdateSignatureByVectorID(ctx context.Context, vectorID types.VectorID, signature types.Signature) error
	UpdateBucketsByVectorID(ctx context.Context, vectorID types.VectorID, bucketsID types.Buckets) error
}
