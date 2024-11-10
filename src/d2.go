package main

import (
	// "fmt"
	"strings"
)

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

func countRunicWords2(input string) int {
  words, runes := parseInput(input)
	wordList := strings.Split(words, ",")
  revRunes := Reverse(runes)

  seen := make(map[int]struct{})
  seenList := make([]string, 0)
  seenListRev := make([]string, 0)

  for _, word := range wordList {
    // fmt.Printf("\n=====\nThe word is '%s'\n", word)
    for i := 0; i + len(word) - 1 < len(runes); i += 1 {
      // fmt.Printf("%s // %s\n", runes[i: i + len(word)], revRunes[i: i + len(word)])
      if runes[i:i + len(word)] == word {
        // fmt.Println("It's a match!")
        seenList = append(seenList, runes[i:i + len(word)])
        for j := i; j < i + len(word); j++ {
          _, ok := seen[j]
          if !ok {
            seen[j] = struct{}{}
          }
        }
      }

      if revRunes[i:i + len(word)] == word {
        // fmt.Println("It's a reverse match!")
        revIndex := len(runes) - i
        seenList = append(seenListRev, runes[revIndex - len(word):revIndex])
        for j := i; j < i + len(word); j++ {
          revIndex = len(runes) - 1 - j
          _, ok := seen[revIndex]
          if !ok {
            seen[revIndex] = struct{}{}
          }
        }
      }
    }
  }

  // fmt.Println(seenList)
  // fmt.Println(seenListRev)

	return len(seen)
}

func parseInput(input string) (words string, text string) {
  splitres := strings.SplitN(input, "\n", 3)
  words = splitres[0]
  words, _ = strings.CutPrefix(words, "WORDS:")

  text = splitres[2]
  return
}
