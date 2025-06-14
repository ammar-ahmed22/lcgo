# 3. Longest Substring Without Repeating Characters

## Problem

Given a string `s`, find the length of the **longest** **substring** without duplicate characters.

**Example 1:**

```
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

```

**Example 2:**

```
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

```

**Example 3:**

```
Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

```

**Constraints:**

- `0 <= s.length <= 5 * 104`
- `s` consists of English letters, digits, symbols and spaces.

## Approach
To solve this problem, we can use a sliding window approach.

The idea is that we keep increasing our window as long as the characters are unique, once we hit a non-unique character, we decrease our window size from the other side (effectively sliding it) until our window is unique again.

In order to do this, we can use a map of character counts with all values initialized to zero. As in, we create hash map of character's as the key's and the integers as the values. We iterate over the characters in the string and set each character in the map to zero.

After this, we initialize our left and right pointers for the window to zero as well as our answer value to zero (we'll keep track of the max length we see in this value). We iterate while the right pointer is less than the length of the string. On each iteration, we check if the count in the map for the character at the right pointer is zero, if it is, we increase our window size and increment the count value. If it is not zero, we want to decrease our window size and slide it by increment the left pointer and decrementing the count at the left pointer because it is being removed from our substring. 

After these two checks, we update our max substring length with `r - l`.

## Complexity
### Time: `O(n)`
We iterate over the string once.

### Space: `O(n)`
We create a hashmap containing all the characters of the string.

## Solution

```go
func lengthOfLongestSubstring(s string) int {
	count := make(map[byte]int)
	for _, c := range s {
		count[byte(c)] = 0
	}

	var (
		l, r, res int
	)
	for r < len(s) {
		rCount := count[s[r]]
		if rCount == 0 {
			count[s[r]]++
			r++
		} else {
			count[s[l]]--
			l++
		}
		res = max(res, r-l)
	}
	return res
}

```
