package interfaces

import (
	"context"

	"github.com/agtabesh/lsh/types"
)

type Store interface {
	GetSignatureByVectorID(ctx context.Context, vectorID types.VectorID) (types.Signature, error)
	GetVectorsIDInBucket(ctx context.Context, bucketID types.Bucket) (chan types.VectorID, error)
	UpdateSignatureByVectorID(ctx context.Context, vectorID types.VectorID, signature types.Signature) error
	UpdateBucketsByVectorID(ctx context.Context, vectorID types.VectorID, bucketsID types.Buckets) error
}
