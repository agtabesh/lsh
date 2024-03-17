package types

type VectorID string
type VectorsID []VectorID

func (v VectorID) String() string {
	return string(v)
}

func (vsID *VectorsID) Remove(vectorID VectorID) {
	for i, vID := range *vsID {
		if vectorID == vID {
			*vsID = append((*vsID)[:i], (*vsID)[i+1:]...)
			break
		}
	}
}
