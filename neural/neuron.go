package neural

type Neuron struct {
	Weights []float64
	Bias    float64
	Lrate   float64
	Delta   float64
	Value   float64
}

func NewNeuron() *Neuron {
	return &Neuron{}
}
