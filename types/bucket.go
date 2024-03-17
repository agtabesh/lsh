package types

type Bucket uint64
type Buckets []Bucket

func (bs Buckets) Contains(b Bucket) bool {
	for _, bucketID := range bs {
		if bucketID == b {
			return true
		}
	}
	return false
}

func (bs Buckets) Diff(bs2 Buckets) Buckets {
	diff := make(Buckets, 0)
	for _, b := range bs {
		if bs2.Contains(b) {
			continue
		}
		diff = append(diff, b)
	}
	return diff
}
