# 242. Valid Anagram

## Problem

Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

**Example 1:**

**Input:** s = "anagram", t = "nagaram"

**Output:** true

**Example 2:**

**Input:** s = "rat", t = "car"

**Output:** false

**Constraints:**

- `1 <= s.length, t.length <= 5 * 104`
- `s` and `t` consist of lowercase English letters.

**Follow up:** What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

## Solution Notes
- We can solve this using a frequency count array
- The idea is that `s` and `t` are anagrams if they have the same frequency of letters
    + Since `s` and `t` are both lowercase English letters, there are only 26 possibilites of letters
    + Therefore, we can create frequency arrays for both where each index maps to each letter of the alphabet
- Create two `int` arrays, `freqS`, `freqT` each of length 26, initialized to zero values
- Iterate over the string `s`
    + For each letter, convert it to it's `int` representation and normalize between 0 and 25
    + Increment the `freqS` array at that index
- Do the same for the string `t`
- Iterate over the `freqS` and `freqT` arrays together and ensure each values matches
    + If there is any mismatch, return `false`
    + Otherwise, if we reach the end of the arrays, -> return `true`

### Further Optimizations
- We can do an initial check to see if `s` and `t` are the same length, if they are not -> return `false`
- Since they are the same length, the frequency count can be done together by iterating over the strings at the same time

### Follow-Up
- If the inputs contained Unicode characters instead of lowercase English letters, we could do the same frequency count but with a hashmap instead to account for a large number of values

### Complexity
#### Time: `O(n)`
- We iterate over the input array only once 
- The second iteration is constant over 26 letters -> no effect

#### Space: `O(1)`
- The only extra space we create is for the frequency arrays which both have constant space
- If we were to go the hashmap route, it would be `O(n)`

## Solution

```go
func isAnagram(s string, t string) bool { 
	if len(s) != len(t) {
		return false
	}

	var freqS [26]int
	var freqT [26]int
	n := len(s)
	for i := range n {
		freqS[s[i] - 'a']++
		freqT[t[i] - 'a']++
	}

	for i := range 26 {
		if freqS[i] != freqT[i] {
			return false
		}
	}

	return true
}
```
