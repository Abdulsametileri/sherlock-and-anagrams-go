package main

import (
	"fmt"
)


func isAnagram(s1, s2 string) bool {
	freq := make([]int, 256)
	for _, c := range s1 {
		freq[c]++
	}

	for _, c := range s2 {
		freq[c]--
	}

	r := 0
	for _, f := range freq {
		if f > 0 {
			r += f
		}
	}

	return r == 0
}

func main() {
	s := "ifailuhkqqhucpoltgtyovarjsnrbfpvmupwjjjfiwwhrlkpekxxnebfrwibylcvkfealgonjkzwlyfhhkefuvgndgdnbelgruel"

	expectedOutput := 399
	_ = expectedOutput

	hashTable := make(map[int][]string, 0)

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			key := s[i : j + 1]
			hash := len(key)
			hashTable[hash] = append(hashTable[hash], key)
		}
	}

	anagramCount := 0

	for _, list := range hashTable {
		for i := 0; i < len(list); i++ {
			for j := i+1; j < len(list); j++ {
				if isAnagram(list[i], list[j]) {
					anagramCount++
				}
			}
		}
	}
	fmt.Println(anagramCount)



}
