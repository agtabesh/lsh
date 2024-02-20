package types

type Qualitative []VectorID

func (q Qualitative) ToQuantitative() Quantitative {
	r := Quantitative{}
	for _, v := range q {
		r[v] = 1
	}
	return r
}
