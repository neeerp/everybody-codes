package main

import (
	"fmt"
	"strings"
)

func countRunicWords(words string, runes string) int {
	wordList := strings.Split(words, ",")

	result := 0
	for _, word := range wordList {
		for i := 0; i+len(word)-1 < len(runes); i += 1 {
			if runes[i:i+len(word)] == word {
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
		for i := 0; i+len(word)-1 < len(runes); i += 1 {
			if runes[i:i+len(word)] == word {
				seenList = append(seenList, runes[i:i+len(word)])
				for j := i; j < i+len(word); j++ {
					_, ok := seen[j]
					if !ok {
						seen[j] = struct{}{}
					}
				}
			}

			if revRunes[i:i+len(word)] == word {
				revIndex := len(runes) - i
				seenList = append(seenListRev, runes[revIndex-len(word):revIndex])
				for j := i; j < i+len(word); j++ {
					revIndex = len(runes) - 1 - j
					_, ok := seen[revIndex]
					if !ok {
						seen[revIndex] = struct{}{}
					}
				}
			}
		}
	}

	return len(seen)
}

type Pair struct {
	a, b interface{}
}

func countRunicWords3(input string) int {
	words, text := parseInput(input)
	wordList := strings.Split(words, ",")

	// Note: The text is going to be an nxk matrix.
	matrix := strings.Split(text, "\n")
	n, m := len(matrix), len(matrix[0])

	seen := make(map[Pair]struct{})
	fmt.Println(matrix)

	for _, word := range wordList {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				cur := getHorizontalWord(matrix[i], len(word), j)
				if cur == word || Reverse(cur) == word {
					fmt.Printf("i: %d; j: %d\n", i, j)
					fmt.Println(cur)
					for k := 0; k < len(word); k++ {
						if k+j < m {
							seen[Pair{i, k + j}] = struct{}{}
						} else {
							seen[Pair{i, k + j - m}] = struct{}{}
						}
						fmt.Println(seen)
					}
				}
			}
		}

		for j := 0; j < m; j++ {
			for i := 0; i < n; i++ {
				cur := getVerticalWord(matrix, len(word), i, j)
				// fmt.Printf("j: %d; i: %d\n", j, i)
				if cur == word || Reverse(cur) == word {
					// fmt.Printf("i: %d; j: %d\n", i, j)
					fmt.Println(cur)
					count := 0
					k := i
					for count != len(word) {
						seen[Pair{k, j}] = struct{}{}
						k += 1
						if k == m {
							k = 0
						}
						count += 1
					}
				}
			}
		}
	}
	fmt.Println(seen)

	return len(seen)
}

func getHorizontalWord(text string, length int, index int) string {
	n := len(text)
	if index+length < n {
		return text[index : index+length]
	}

	remaining := length - (n - index)
	return text[index:] + text[:remaining]
}

func getVerticalWord(matrix []string, length int, row int, col int) string {
	m := len(matrix)

	if row+length > m {
		return ""
	}

	result := ""
	i := row
	for len(result) != length {
		result = result + string(matrix[i][col])
		i += 1
		if i == m {
			i = 0
		}
	}

	return result
}

func parseInput(input string) (words string, text string) {
	splitres := strings.SplitN(input, "\n", 3)
	words = splitres[0]
	words, _ = strings.CutPrefix(words, "WORDS:")

	text = splitres[2]
	return
}
