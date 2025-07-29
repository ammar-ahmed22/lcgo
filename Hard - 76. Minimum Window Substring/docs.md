# 76. Minimum Window Substring

## Problem

Given two strings `s` and `t` of lengths `m` and `n` respectively, return _the **minimum window**_ **_substring_** _of_ `s` _such that every character in_ `t` _( **including duplicates**) is included in the window_. If there is no such substring, return _the empty string_ `""`.

The testcases will be generated such that the answer is **unique**.

**Example 1:**

```
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

```

**Example 2:**

```
Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.

```

**Example 3:**

```
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.

```

**Constraints:**

- `m == s.length`
- `n == t.length`
- `1 <= m, n <= 105`
- `s` and `t` consist of uppercase and lowercase English letters.

**Follow up:** Could you find an algorithm that runs in `O(m + n)` time?

## Approach
To solve this problem, we can use a sliding window alongside a frequency map.

To start, we create the frequency map for the characters in `t`. This is because we are looking for when a window substring of `s` that contains all of the letters from `t` including duplicates. Once the frequency map is created, we initialize our window. Since the minimum length of the window is the length of `t`, we start our window size at that with `l = 0` and `r = len(t) - 1`. 

Now, we want to process our first window before we start iterating. The processing will include subtracting the value for each character in the window that exists in the frequency map. So, using example 1, our first window will be `"ADO"` and our frequency map will look like: `{ "A": 1, "B": 1, "C": 1 }`. Since our first window contains an `A`, we subtract it from the frequency map: `{ "A": 0, "B": 1, "C": 1 }`. 

We will follow this same pattern whenever we add a new character to the window and the opposite when removing a character from the window. From this, we can see that whenever the frequency map contains only values less than or equal to zero, we have potential answer.

Once the first window is processed, we can start iterating. We'll iterate while the right pointer is less than the length of `s`. On each iteration, we first check if the current window substring is valid by checking if all values in the frequency map are less than or equal to zero.

If the substring is valid, we update our answer with this potential answer (if it is smaller than the current answer). Then, if our current window is greater than our minimum window length, we decrease our window size from the left. We also must ensure to update our frequency map **before** updating the left pointer because that character is "leaving" the window. If our window length is already at the minimum length, we want to slide our window by also moving the right pointer. The frequency map should be updated **after** the pointer is moved because the character is being added to the window. We should also check that our right pointer is still in bounds before doing this update.

If the substring is not valid, we want to increase our window size from the right, following the same steps as above in regard to updating the frequency map.

After the iteration is complete, return the answer.

Also, ensure to check for the edge case at the top of the function where `t` is greater than `s` in length, in this case there is no possible string so return an empty string.

## Complexity
### Time: `O(m + n)`
We iterate over the string `t` once to create the frequency map, `O(n)`. After that our sliding window approach iterates over the string `s` once as well, `O(m)`.

### Space: `O(n)`
We create the frequency map that will have at most `n` elements.

## Solution

```go
func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	ans := ""
	freq := make(map[byte]int)
	for _, char := range t {
		freq[byte(char)]++
	}

	l, r := 0, len(t)-1

	for i := range len(t) {
		if _, ok := freq[byte(s[i])]; ok {
			freq[byte(s[i])]--
		}
	}

	minWindowLen := len(t)

	for r < len(s) {
		isValid := true
		for _, val := range freq {
			if val > 0 {
				isValid = false
				break
			}
		}

		if isValid {
			windowLen := (r - l + 1)
			if ans == "" || windowLen < len(ans) {
				ans = s[l : r+1]
			}

			if _, ok := freq[byte(s[l])]; ok {
				freq[byte(s[l])]++
			}
			l++

			if windowLen == minWindowLen {
				r++
				if r < len(s) {
					if _, ok := freq[byte(s[r])]; ok {
						freq[byte(s[r])]--
					}
				}
			}
		} else {
			r++
			if r < len(s) {
				if _, ok := freq[byte(s[r])]; ok {
					freq[byte(s[r])]--
				}
			}
		}
	}

	return ans
}

```
