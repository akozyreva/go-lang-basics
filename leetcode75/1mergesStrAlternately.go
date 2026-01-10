// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
// If a string is longer than the other, append the additional letters onto the end of the merged string.
package main

import "fmt"

func mergeAlternately(word1 string, word2 string) {
	word1Len := len(word1)
	word2Len := len(word2)
	var mergedWord string
	var index int
	if word1Len > word2Len || word1Len == word2Len {
		index = word1Len
	} else {
		index = word2Len
	}
	for i := 0; i < index; i++ {
		if i+1 <= word1Len {
			mergedWord += string(word1[i])
		}
		if i+1 <= word2Len {
			mergedWord += string(word2[i])
		}

	}
	fmt.Println(mergedWord)
}

func mergeAlternatelySolution(word1 string, word2 string) {
	var result []byte // it's slice

	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			result = append(result, word1[i])
		}
		if i < len(word2) {
			result = append(result, word2[i])
		}
	}

	fmt.Println(string(result))
}

func main() {
	mergeAlternately("abc", "pqr")
	mergeAlternately("abcd", "pq")
	mergeAlternatelySolution("abc", "pqr")
}
