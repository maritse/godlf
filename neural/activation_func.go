package neural

import (
	"math"
)

func sigmoid(value float64) float64 {
	return 1.0 / (1.0 + math.Exp(-value))
}

func sigmoidDerivative(value float64) float64 {
	// calculation -> https://math.stackexchange.com/questions/78575/derivative-of-sigmoid-function-sigma-x-frac11e-x
	return sigmoid(value) * (1.0 - sigmoid(value))
}
