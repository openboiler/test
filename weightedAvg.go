package main

import (
	"fmt"
)

type WeightedAvgElement struct {
	Weight float32
	Value  float32
}

func WeightedAvg(elements []WeightedAvgElement) float32 {
	var numerator float32
	var denomominator float32
	for _, e := range elements {
		numerator += (e.Value * e.Weight)
		denomominator += e.Weight
	}
	return numerator / denomominator
}

func avg(elements []WeightedAvgElement) float32 {
	var sum float32
	for _, e := range elements {
		sum += e.Value
	}
	return sum / float32(len(elements))
}

func main() {
	elements := make([]WeightedAvgElement, 5, 5)
	elements[0].Value = 5
	elements[0].Weight = 1
	elements[1].Value = 5
	elements[1].Weight = 2
	elements[2].Value = 6
	elements[2].Weight = 3
	elements[3].Value = 6
	elements[3].Weight = 4
	elements[4].Value = 6
	elements[4].Weight = 50
	fmt.Printf("%f vs %f\n", WeightedAvg(elements), avg(elements))
}
