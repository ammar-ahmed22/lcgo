package main

import (
	. "group-anagrams/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)

	for _, str := range strs {
		var freq [26]int
		for _, ch := range str {
			freq[ch - 'a']++
		}

		if words, exists := group[freq]; exists {
			group[freq] = append(words, str)
		} else {
			group[freq] = []string{str}
		}
	}

	var result [][]string
	for _, value := range group {
		result = append(result, value)
	}
	return result 
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = [][]string

var testCases = []*TestCase[ReturnType]{
	NewTestCase([][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}).WithArgs([]string{"eat", "tea", "tan", "ate", "nat", "bat"}),
	NewTestCase([][]string{{""}}).WithArgs([]string{""}),
	NewTestCase([][]string{{"a"}}).WithArgs([]string{"a"}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return groupAnagrams(args[0].([]string))
	})
}
