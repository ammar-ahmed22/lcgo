# 567. Permutation in String

## Problem

Given two strings `s1` and `s2`, return `true` if `s2` contains a permutation of `s1`, or `false` otherwise.

In other words, return `true` if one of `s1`'s permutations is the substring of `s2`.

**Example 1:**

```
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").

```

**Example 2:**

```
Input: s1 = "ab", s2 = "eidboaoo"
Output: false

```

**Constraints:**

- `1 <= s1.length, s2.length <= 104`
- `s1` and `s2` consist of lowercase English letters.

## Approach
To solve this problem, we can use a sliding window approach alongside a frequency map. 

The conditions for returning true are when there is a substring in `s2` that has the same frequency of letters as the string `s1`, that means that `s2` contains a permutation of that string. This also means that we should only be checking substrings of the same length of `s1`.

Therefore, the basic premise of solving this problem includes sliding a window of length `s1` across the string `s2` and checking if the frequency of letters is the same as that of `s1`.

To do this efficiently, we can use a frequency map that we update dynamically as we slide the window. We can also employ the fact that both `s1` and `s2` consist only of lowercase English letters to further optimize our map by using a constant length array of 26 values to track the frequency (there are only 26 possible characters).

A slightly less efficient but easier to understand algorithm would be using two distinct maps, one for the string `s1` and one for our current operating window. On each iteration, we check if the maps are the same, if they are return true. Otherwise, if we get to the end, return false. However, we can optimize this even further by using a single map. We start by adding the frequency of `s1` to the map. Then, for each letter that is in our window, we subtract from the frequency map. Therefore, if the map ever becomes all zeroes, we know we have reached a valid substring.

For the actual implementation of this algorithm, we can start by doing an edge case check, if the length of `s1` is greater than `s2`, there are no possible permutations in the string so we can return false early. Next, we can initialize our constant length array frequency map. After this, we can start by adding the characters in `s1` to the map. Then, we initialize our left and right pointers; left will be at zero to start and right will be at the length of `s1` so that our window is always the same length. From this, we can update our frequency map with the letters in the current window by subtracting them from the frequency map (because they are "entering" the window). Now we are ready to start sliding our window. We iterate while `right` is less than the length of `s2`. 

On each iteration, we start by checking if our map is all zeros. If it is, return true. Otherwise, we need to slide our window. Since the value at the right pointer is "entering" our window, we decrement that character from the map. Our left pointer value is "leaving" the window so we increment that character in the map. 

After the iteration is complete, we need to do a final check for the last substring by checking if the map is all zeros, if it is return true, otherwise return falsewe increment that character in the map. 

After the iteration is complete, we need to do a final check for the last substring by checking if the map is all zeros, if it is return true, otherwise return false.

## Complexity
### Time: `O(n + m)`
Where `m` is the length of `s1` and `n` is the length of `s2`. We iterate over `s1` once to add it to the frequency map. Then we iterate over `s2` once as well while sliding our window. 

### Space: `O(1)`
The only extra space created is the frequency map which is constant length of 26; `O(26) = O(1)`.

## Solution

```go
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	// Populate frequency map with s1 values
	var freqMap [26]int
	for _, char := range s1 {
		freqMap[char-'a']++
	}

	l := 0
	r := len(s1)

	// Populate frequency map with chars in s2 from l to r
	for i := range r {
		// Entering "window", subtract from map
		freqMap[s2[i]-'a']--
	}

	for r < len(s2) {
		// Check if freqMap is all zeros
		hasNonZero := false
		for _, f := range freqMap {
			if f != 0 {
				hasNonZero = true
				break
			}
		}
		if !hasNonZero {
			return true
		}

		// Leaving window
		freqMap[s2[l]-'a']++
		// Entering window
		freqMap[s2[r]-'a']--
		l++
		r++
	}

	for _, f := range freqMap {
		if f != 0 {
			return false
		}
	}
	return true
}

```
