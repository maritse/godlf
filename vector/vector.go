package vector

type Vector []float64

func (v Vector) sum() (sumResult float64) {
	for _, i := range v {
		sumResult += i
	}
	return sumResult
}

func (v Vector) isEqual(w Vector) bool {
	if len(v) != len(w) {
		return false
	}
	for index, _ := range v {
		if v[index] != w[index] {
			return false
		}
	}
	return true
}
