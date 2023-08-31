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
		{[]string{"aya", "ye", "woo", "ma"}, 4}, // Added a test case based on the provided WakeUp words
		{[]string{}, 0},                         // Test for empty input
		{make([]string, 101), 0},                // Test for input exceeding the length limit
	}

	for _, test := range tests {
		result := listenBabbling(test.babbling)
		if result != test.expected {
			t.Errorf("For babbling %v, expected %d, but got %d", test.babbling, test.expected, result)
		}
	}
}
