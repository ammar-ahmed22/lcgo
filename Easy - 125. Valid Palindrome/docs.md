# 125. Valid Palindrome

## Problem

A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string `s`, return `true` _if it is a **palindrome**, or_ `false` _otherwise_.

**Example 1:**

```
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

```

**Example 2:**

```
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

```

**Example 3:**

```
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

```

**Constraints:**

- `1 <= s.length <= 2 * 105`
- `s` consists only of printable ASCII characters.

## Approach
We can use a two pointer approach to solve this problem. In order to check if the string is the same backwards and forwards we essentially want to check if the string has the same letters from both sides, hence the two pointers.

We can have a left and right pointer which we used to iterate while the left pointer is less than the right pointer. On each iteration, we need to do an important check of essentially ignoring any characters that are not alphanumeric. If the character at the left pointer is not alphanumeric, we increment it and continue. If the character on the right is not an alphanumeric, we decrement the right pointer and continue. After these checks, we can simply check if the right and left characters are equal, if not we can return `false` early. Otherwise, we can increment the left and decrement the right and continue on. If we complete the iteration, that means the string is the same from both sides, i.e. it is a palindrome and we can return true. We can also convert the string to lowercase at the start to make our alphanumeric check a little simpler but it's not necessary.

For my solution, I wrote a helper function to do the `isAlphanumeric` check but this could also be inlined.

## Solution

```go
import "strings"
func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
func isPalindrome(s string) bool { 
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1
	for l < r {
		left := s[l]
		right := s[r]
		if !isAlphaNumeric(left) {
			l++
			continue
		}

		if !isAlphaNumeric(right) {
			r--
			continue
		}
		
		if left != right {
			return false
		}
		l++
		r--
	}
	return true
}
```
