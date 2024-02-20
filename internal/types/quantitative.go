package types

type Quantitative map[VectorID]float64

func (q Quantitative) ToQuantitative() Quantitative {
	return q
}
