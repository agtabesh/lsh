package hash_family

import (
	"hash"

	"github.com/agtabesh/lsh/types"
	"github.com/pierrec/xxHash/xxHash64"
)

type XXHASH64HashFamily struct {
	hasher []hash.Hash64
	count  int
}

// NewXXHASH64HashFamily initializes a new HashFamily with the specified number of hash functions.
func NewXXHASH64HashFamily(count int) XXHASH64HashFamily {
	hasher := make([]hash.Hash64, count)
	for i := 0; i < count; i += 1 {
		hasher[i] = xxHash64.New(uint64(i)) // Initialize each hash function with a unique seed
	}
	return XXHASH64HashFamily{
		hasher: hasher,
		count:  count,
	}
}

// Hash computes the hash values of a given string using all hash functions in the family.
func (hf XXHASH64HashFamily) Hash(s string) types.Signature {
	hashes := make(types.Signature, hf.count)
	for i := 0; i < hf.count; i += 1 {
		b := []byte(s)
		hf.hasher[i].Reset()
		hf.hasher[i].Write(b)
		h := hf.hasher[i].Sum64()
		hashes[i] = types.SignatureEntry(h)
	}
	return hashes
}

// MinHash computes the min hash values of a given vector using all hash functions in the family.
func (hf XXHASH64HashFamily) MinHash(vector types.Vector) types.Signature {
	var signature types.Signature
	for k := range vector {
		newSignature := hf.Hash(k.String())
		signature = signature.FindMin(newSignature)
	}
	return signature
}
