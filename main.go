package main

import "fmt"

func main() {

	inputs := []float32{1, 2, 3, 2.5}

	weights := [][]float32{
		{0.2, 0.8, -0.5, 1.0},
		{0.5, -0.91, 0.26, -0.5},
		{-0.26, -0.27, 0.17, 0.87},
	}

	biases := []float32{2.0, 3.0, 0.5}


	layerOutputs := make([]float32, 3)

for i := range weights{
neuronWeights := weights[i]
neuroBias := biases[i]

var neuronOutput float32 = 0
for j := range inputs{
	neuronOutput += inputs[j] * neuronWeights[j]
}
neuronOutput += neuroBias
layerOutputs[i] = neuronOutput
}

	fmt.Println(layerOutputs)

}
