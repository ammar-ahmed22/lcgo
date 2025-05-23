# 49. Group Anagrams

## Problem

Given an array of strings `strs`, group the anagrams together. You can return the answer in **any order**.

**Example 1:**

**Input:** strs = \["eat","tea","tan","ate","nat","bat"\]

**Output:**\[\["bat"\],\["nat","tan"\],\["ate","eat","tea"\]\]

**Explanation:**

- There is no string in strs that can be rearranged to form `"bat"`.
- The strings `"nat"` and `"tan"` are anagrams as they can be rearranged to form each other.
- The strings `"ate"`, `"eat"`, and `"tea"` are anagrams as they can be rearranged to form each other.

**Example 2:**

**Input:** strs = \[""\]

**Output:**\[\[""\]\]

**Example 3:**

**Input:** strs = \["a"\]

**Output:**\[\["a"\]\]

**Constraints:**

- `1 <= strs.length <= 104`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

## Solution Notes
- A little more complex but similar to the [Valid Anagram's question](./Easy%20-%20242.%20Valid%20Anagram/docs.md)
- We still will need to do the same thing in terms of counting frequencies of letters, however, we now want to track multiple anagrams
- We can use a hash map to store the frequency arrays alongside the words that correspond to that
- Create a hash map with the frequency array as the key and a slice of strings as the value (`map[26[int]][]string`)
- Iterate over the `strs`
    + For each word, create the frequency array
    + If that frequency array already exists in the hash map, add the word to it's corresponding string slice
    + If it doesn't, add a new slice with that word
- Iterate over the hash map and create the result

### Complexity
#### Time: `O(n)`
- We iterate over the input array once -> `O(n)`
- We also iterate over the hash map once, which could potentially contain `n` elements if none are anagrams -> `O(n)`
- Therefore, `O(n) + O(n) = O(n)`

#### Space: `O(n)`
- The hash map can potentially contain `n` entries

## Solution

```go
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
```
