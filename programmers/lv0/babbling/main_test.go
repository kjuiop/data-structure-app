package main

import "testing"

func TestWakeUp(t *testing.T) {
	baby := WakeUp()

	expectedWords := []string{"aya", "ye", "woo", "ma"}
	if len(baby.words) != len(expectedWords) {
		t.Errorf("Expected %d words, but got %d", len(expectedWords), len(baby.words))
		return
	}

	for i, expectedWord := range expectedWords {
		if baby.words[i] != expectedWord {
			t.Errorf("Expected word '%s', but got '%s'", expectedWord, baby.words[i])
		}
	}
}

func TestListenBabbling(t *testing.T) {
	tests := []struct {
		babbling []string
		expected int
	}{
		{[]string{"ayaye", "ayaye", "uuuma", "ye", "yemawoo", "ayaa"}, 4},
		{[]string{"aya", "yee", "u", "maa", "wyeoo"}, 1},
		{[]string{"ayaye", "uuuma", "ye", "yemawoo", "ayaa"}, 3},
	}

	for _, test := range tests {
		result := listenBabbling(test.babbling)
		if result != test.expected {
			t.Errorf("For babbling %v, expected %d, but got %d", test.babbling, test.expected, result)
		}
	}
}

func TestSpeakWords(t *testing.T) {
	baby := WakeUp()

	tests := []struct {
		babbling []string
		expected int
	}{
		{[]string{"ayaye", "ayaye", "uuuma", "ye", "yemawoo", "ayaa"}, 4},
		{[]string{"aya", "yee", "u", "maa", "wyeoo"}, 1},
		{[]string{"ayaye", "uuuma", "ye", "yemawoo", "ayaa"}, 3},
	}

	for _, test := range tests {
		result := baby.speakWords(test.babbling)
		if result != test.expected {
			t.Errorf("For babbling %v, expected %d, but got %d", test.babbling, test.expected, result)
		}
	}
}

func TestAnalyzeBabyBabble(t *testing.T) {
	baby := WakeUp()

	tests := []struct {
		babble   string
		expected bool
	}{
		{"aya", true},
		{"ye", true},
		{"wooma", true},
		{"ayawooma", true},
		{"ayayewoomama", false},
		{"ayayewooma", true},
		{"ayaayayewooma", false},
	}

	for _, test := range tests {
		result := baby.analyzeBabyBabble(test.babble, makeTargetWordsMap(baby.words))
		if result != test.expected {
			t.Errorf("For babble '%s', expected %t, but got %t", test.babble, test.expected, result)
		}
	}
}

func makeTargetWordsMap(words []string) map[string]bool {
	targetWords := make(map[string]bool)
	for _, word := range words {
		targetWords[word] = true
	}
	return targetWords
}
