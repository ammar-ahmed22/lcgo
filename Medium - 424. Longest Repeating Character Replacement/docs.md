# 424. Longest Repeating Character Replacement

## Problem

You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return _the length of the longest substring containing the same letter you can get after performing the above operations_.

**Example 1:**

```
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

```

**Example 2:**

```
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.
```

**Constraints:**

- `1 <= s.length <= 105`
- `s` consists of only uppercase English letters.
- `0 <= k <= s.length`

## Approach
To solve this problem efficiently, we can start with the basic idea. For each substring, we want to replace the other characters with the most frequent character in the substring to maximize our length. The length of the substring minus the frequency of the most frequent one will be the number of operations required to make it valid.

For example, if we have the substring, `"AABA"`, we have three "A"'s and one "B". Therefore, for this length of string, we have to do `4 - 3 = 1` operations to get it to be a valid string.

With this information in hand, we can employ a sliding window technique alongside tracking the max frequency we see.

On each iteration, we check if the length of the current substring minus the maximum frequency we have seen is less than or equal to `k`, if it is, we can update our max length value. If our check is not true, we slide the window from the left side and continue. We only need to track the maximum frequency we have seen regardless of what sliding the window does because we want to maximize our length. Therefore, if the sliding/shortening of the window results in a smaller frequency value, we don't care because this will only decrease our max length.

We can track the frequency using a fixed length array of 26 values because the string consists only of uppercase English letters.

The algorithm goes as follows:
- Initialize our left pointer, max frequency and max length variables
- Iterate with the right pointer always moving forward
- First, we increment the frequency of the character at the right pointer and update the max frequency value because we incremented a frequency value
- Next, we iterate while the length of the substring minus the max frequency is greater than `k` and move the left pointer over decrementing the frequency value at the left pointer. This essentially shrinks our window until we are at a valid substring.
- Since we are now at a valid substring, we can update our max length value

Once again, the reason we don't update the max frequency value while we are shrinking our window is because we are only decreasing frequency values there. Since we are maximizing for length, it's possible that the max frequency value is not the max frequency of that specific substring but that would only cause us to overestimate our max length which is what we want. In other words, we might miss some of the smaller substrings in our calculations but that's a good thing since we only care about the largest ones.


## Complexity
### Time: `O(n)`
We iterate over the string only once.

### Space: `O(1)`
The only extra space created is the frequency array which is a constant length of 26 

## Solution

```go
func characterReplacement(s string, k int) int {
	var freq [26]int
	left, maxF, maxLen := 0, 0, 0

	for right := range len(s) {
		freq[s[right]-'A']++
		maxF = max(maxF, freq[s[right]-'A'])

		for (right-left+1)-maxF > k {
			freq[s[left]-'A']--
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

```
