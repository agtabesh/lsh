package types

import "sort"

type Similarities map[VectorID]float64

func (s Similarities) Top(count int) []VectorID {
	if len(s) == 0 {
		return []VectorID{}
	}
	if count > len(s) {
		count = len(s)
	}
	similarVectorsID := make([]VectorID, 0, len(s))
	for key := range s {
		similarVectorsID = append(similarVectorsID, key)
	}
	sort.SliceStable(similarVectorsID, func(i, j int) bool {
		return s[similarVectorsID[i]] > s[similarVectorsID[j]]
	})
	return similarVectorsID[:count]
}
