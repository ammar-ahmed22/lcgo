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

## Approach 
To solve this, we can use a frequency count array. Since the strings consist only of lowercase English letters, there are only 26 possible letter that we need to track. Therefore, we can create constant size `int` arrays initialized to zero with size 26. From this, we can iterate over the strings together (we will have already checked they are the same length) and increment the corresponding index for the letter in each array. Finally, we can check if the arrays have the same values, if they don't, we return `false` early. Otherwise we return `true`.

Another optimization that I'm thinking of now (not included in the solution) is that we don't need two arrays actually. We can increment the value for one string, say `s` and decrement for the other, `t`. Then, we can simply check if all the values in the frequency count array are zero.

### Follow-Up
If the inputs contained Unicode characters instead of lowercase English letters, we could do the same frequency count but with a hashmap instead to account for a larger number of values. We could also use a hashmap for the original problem, however a constant size array is just teeny bit faster because the memory is allocated to start with. 

## Complexity
### Time: `O(n)`
- We iterate over the input array only once 
- The second iteration is constant over 26 letters -> no effect

### Space: `O(1)`
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
