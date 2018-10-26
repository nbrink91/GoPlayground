package main

type Transcript struct {
	Utterances []Utterance `json:"utterances"`
}

type Metrics struct {
	UtteranceCount    int     `json:"utteranceCount"`
	AverageConfidence float32 `json:"averageConfidence"`
}

type Utterance struct {
	Text       string  `json:"text"`
	Confidence float32 `json:"confidence"`
	Start      float32 `json:"start"`
	End        float32 `json:"end"`
}

func GetMetrics(transcript Transcript) Metrics {
	length := make(chan int)
	go GetLength(length, transcript.Utterances)

	avgConfidence := make(chan float32)
	go GetAverageConfidence(avgConfidence, transcript.Utterances)

	return Metrics{
		UtteranceCount:    <-length,
		AverageConfidence: <-avgConfidence,
	}
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
