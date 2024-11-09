package main

import "strings"

func countRunicWords(words string, runes string) int {
	wordList := strings.Split(words, ",")

  result := 0
  for _, word := range wordList {
    for i := 0; i + len(word) - 1 < len(runes); i += 1 {
      if runes[i:i + len(word)] == word {
        result += 1
      }
    }

  }

	return result
}
