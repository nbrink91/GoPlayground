package main

import (
	"fmt"
)

type Metrics struct {
	utterenceCount    int
	averageConfidence float32
}

type Utterance struct {
	id         int32
	confidence float32
}

func main() {
	transcript := GetTranscript()

	length := make(chan int)
	go GetLength(length, transcript)

	avgConfidence := make(chan float32)
	go GetAverageConfidence(avgConfidence, transcript)

	fmt.Println(Metrics{utterenceCount: <-length, averageConfidence: <-avgConfidence})
}

func GetLength(length chan<- int, utterences []Utterance) {
	length <- len(utterences)
}

func GetAverageConfidence(avgConfidence chan<- float32, utterences []Utterance) {
	var total float32 = 0
	for _, v := range utterences {
		total += v.confidence
	}

	avgConfidence <- total / float32(len(utterences))
}

func GetTranscript() []Utterance {
	data := []Utterance{
		Utterance{id: 1, confidence: 1.0},
		Utterance{id: 2, confidence: 0.75},
	}
	return data
}
