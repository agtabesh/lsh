package interfaces

import "github.com/agtabesh/lsh/internal/types"

type HashFamily interface {
	Hash(s string) types.Signature
	MinHash(v types.Vector) types.Signature
}
