package metrics

import "testing"

var utterances = []Utterance{
	Utterance{Start: 0, End: 1, Confidence: 1, Text: "Hello"},
	Utterance{Start: 1, End: 2, Confidence: 0.5, Text: "There"},
}

func TestGetLength(t *testing.T) {
	length := make(chan int)
	go GetLength(length, utterances)
	expected := 2

	result := <-length
	if result != expected {
		t.Errorf("Expected length to be %d, got %d", expected, result)
	}
}

func TestGetAverageConfidence(t *testing.T) {
	avgConfidence := make(chan float32)

	go GetAverageConfidence(avgConfidence, utterances)

	var expected float32 = 0.75

	result := <-avgConfidence

	if result != expected {
		t.Errorf("Expected average confidence of %f, got %f", expected, result)
	}
}
