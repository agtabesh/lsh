package types

type SignatureEntry uint64
type Signature []SignatureEntry
type Signatures []Signature

func (s Signature) Buckets(numOfBands int) Buckets {
	rowsPerBand := len(s) / numOfBands
	buckets := make(Buckets, numOfBands)
	for i := 0; i < numOfBands; i += 1 {
		sum := Bucket(0)
		for j := 0; j < rowsPerBand; j += 1 {
			k := i*rowsPerBand + j
			sum += Bucket(s[k] / SignatureEntry(rowsPerBand))
		}
		buckets[i] = sum
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
