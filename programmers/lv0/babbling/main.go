package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("result : ", listenBabbling([]string{"ayaye", "ayaye", "uuuma", "ye", "yemawoo", "ayaa"}))
	fmt.Println("result : ", listenBabbling([]string{"aya", "yee", "u", "maa", "wyeoo"}))
	fmt.Println("result : ", listenBabbling([]string{"ayaye", "uuuma", "ye", "yemawoo", "ayaa"}))
}

func listenBabbling(babbling []string) int {
	if len(babbling) < 1 || len(babbling) > 100 {
		return 0
	}

	baby := WakeUp()
	words := baby.speakWords(babbling)

	return words
}

type Baby struct {
	words []string
}

func WakeUp() *Baby {
	return &Baby{words: []string{"aya", "ye", "woo", "ma"}}
}

func (b *Baby) speakWords(babbling []string) int {
	if len(babbling) == 0 {
		return 0
	}

	// 값복사
	targetWords := make(map[string]bool)
	for _, word := range b.words {
		targetWords[word] = true
	}

	// 말한 단어의 갯수
	var babySpeakWords []string

	// 옹알이 검사 시작
	for _, babble := range babbling {
		if len(babble) < 1 || len(babble) > 15 {
			continue
		}
		if listen := b.analyzeBabyBabble(babble, targetWords); !listen {
			continue
		}
		// already 추가
		babySpeakWords = append(babySpeakWords, babble)
	}

	return len(babySpeakWords)
}

func (b *Baby) analyzeBabyBabble(babble string, targetWords map[string]bool) bool {

	for word, _ := range targetWords {
		if !strings.HasPrefix(babble, word) {
			continue
		}
		if strings.Count(babble, word) > 1 {
			continue
		}
		split := strings.Replace(babble, word, "", 1)
		if len(split) == 0 {
			return true
		}
		return b.analyzeBabyBabble(split, targetWords)
	}

	return false
}
