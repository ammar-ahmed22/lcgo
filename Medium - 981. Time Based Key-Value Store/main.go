package main

import (
	. "time-based-key-value-store/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
type TimeMap struct {
	timestamps map[string][]int
	values     map[string][]string
}

func Constructor() TimeMap {
	return TimeMap{
		timestamps: make(map[string][]int),
		values:     make(map[string][]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timestamps[key] = append(this.timestamps[key], timestamp)
	this.values[key] = append(this.values[key], value)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	timestamps, exists := this.timestamps[key]
	if !exists {
		return ""
	}
	values, exists := this.values[key]
	if !exists {
		return ""
	}


	i := -1
	l, r := 0, len(timestamps)-1
	for l <= r {
		m := (l + r) / 2
		if timestamps[m] == timestamp {
			return values[m]
		} else if timestamps[m] < timestamp {
			i = m
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if i == -1 {
		return  ""
	}

	return values[i]
}
// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = TimeMap

var testCases = []*TestCase[ReturnType]{
	// Add test cases here
	// Using twoSum as an example:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9),
	// You can also give a name to the test case:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithName("Example Test Case"),
	// You can also provide a custom comparison function:
	// NewTestCase([]int{0, 1}).WithArgs([]int{2, 7, 11, 15}, 9).WithCompareFn(SliceEqualUnordered)
}

func main() {
	// RunTestCases(testCases, func(args ...any) ReturnType {
	// 	return type TimeMap struct { } func Constructor()
	// })
	timeMap := Constructor()
	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)
	Test("Get (ts = 5, set(love, 10), set(love, 20))", "", timeMap.Get("love", 5))
	Test("Get (ts = 10, set(love, 10), set(love, 20))", "high", timeMap.Get("love", 10))
	Test("Get (ts = 15, set(love, 10), set(love, 20))", "high", timeMap.Get("love", 15))
	Test("Get (ts = 20, set(love, 10), set(love, 20))", "low", timeMap.Get("love", 20))
	Test("Get (ts = 25, set(love, 10), set(love, 20))", "low", timeMap.Get("love", 25))
	// timeMap.Set("foo", "bar2", 4)
	// Test("Get (ts = 4, set(bar, 1), set(bar2, 4))", "bar2", timeMap.Get("foo", 4))
	// Test("Get (ts = 5, set(bar, 1), set(bar2, 4))", "bar2", timeMap.Get("foo", 5))
}
