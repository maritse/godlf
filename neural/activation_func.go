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

func ReLU(value float64) float64 {
	if value > 0 {
		return value
	}
	return 0
}

func ReLUDerivative(value float64) float64 {
	if value < 0 {
		return value
	}
	return 0
}

func linear(value float64) float64 {
	return value
}

func linearDerivative(value float64) float64 {
	return 1
}
