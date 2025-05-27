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

## Approach 
This question builds upon the [Valid Anagram's question](https://ammarahmed.ca/easy/valid-anagram). We'll use the same approach of counting letter frequencies using a constant size array of 26 values, however, now we want to track multiple anagrams.
We can use a hashmap to store the frequency arrays as the keys and an array of strings as it's values.
We start by iterating over the strings. For each word, we create it's frequency array. Then, we check if that specific frequency array exists in the hashmap. If it does, we add that string to it's value array. If it doesn't, we create a new entry in the hashmap with that frequency array as the key and an array of strings with that string as it's first value as the hashmap value. To create the result, we simply iterate over the hashmap and combine all the values.

## Complexity
### Time: `O(n)`
We iterate over the input array once (`O(n)`) and we also iterate over the hash map once, which could potentially have `n` values (`O(n)`)

### Space: `O(n)`
The hash map can potentially contain `n` entries

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
