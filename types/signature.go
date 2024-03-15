package types

type SignatureEntry uint64
type Signature []SignatureEntry
type Signatures []Signature

// this function generate merkle tree of the signature to create buckets
func (s Signature) Buckets() Buckets {
	size := len(s) - 1
	buckets := make(Buckets, size)
	i := 0
	for i < size {
		rowsPerBand := len(s) / (i + 1)
		numOfBands := len(s) / rowsPerBand
		for k := 0; k < numOfBands; k += 1 {
			sum := Bucket(0)
			for j := 0; j < rowsPerBand; j += 1 {
				m := k*rowsPerBand + j
				sum += Bucket(s[m] / SignatureEntry(rowsPerBand))
			}
			buckets[i] = sum
			i += 1
			if i == size {
				break
			}
		}
	}
	return buckets
}

func (s Signature) IsEmpty() bool {
	return len(s) == 0
}

func (s Signature) FindMin(o Signature) Signature {
	if s.IsEmpty() && !o.IsEmpty() {
		return o
	}
	if !s.IsEmpty() && o.IsEmpty() {
		return s
	}
	signature := make(Signature, len(s))
	for i := 0; i < len(s); i++ {
		signature[i] = min(s[i], o[i])
	}
	return signature
}
