package main

type Metrics struct {
	UtteranceCount    int     `json:"utterenceCount"`
	AverageConfidence float32 `json:"averageConfidence"`
}

type Utterance struct {
	ID         int32
	Confidence float32
}

func GetMetrics() Metrics {
	transcript := GetTranscript()

	length := make(chan int)
	go GetLength(length, transcript)

	avgConfidence := make(chan float32)
	go GetAverageConfidence(avgConfidence, transcript)

	return Metrics{UtteranceCount: <-length, AverageConfidence: <-avgConfidence}
}

func GetLength(length chan<- int, utterances []Utterance) {
	length <- len(utterances)
}

func GetAverageConfidence(avgConfidence chan<- float32, utterances []Utterance) {
	var total float32 = 0
	for _, utterance := range utterances {
		total += utterance.Confidence
	}

	avgConfidence <- total / float32(len(utterances))
}

func GetTranscript() []Utterance {
	return []Utterance{
		Utterance{ID: 1, Confidence: 1.0},
		Utterance{ID: 2, Confidence: 0.75},
	}
}
