package main

import (
	. "car-fleet/testutils"
)

// <-- DO NOT REMOVE: PROBLEM START -->
import "sort"
func carFleet(target int, position []int, speed []int) int {
	// Combine position and speed
	combined := make([][]int, len(position))
	for i := range position {
		combined[i] = []int{position[i], speed[i]}
	}

	// Sort combined by position
	sort.Slice(combined, func(i, j int) bool {
		return combined[i][0] < combined[j][0]
	})


	var stack []float32
	for i := len(combined) - 1; i >= 0; i-- {
		pos := combined[i][0]
		speed := combined[i][1]
		timeToTarget := float32(target - pos) / float32(speed)
		if len(stack) > 0 && timeToTarget <= stack[len(stack) - 1] {
			continue
		}
		stack = append(stack, timeToTarget)
	}

	return len(stack)
}

// <-- DO NOT REMOVE: PROBLEM END -->

type ReturnType = int

var testCases = []*TestCase[ReturnType]{
	NewTestCase(3).WithArgs(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}),
	NewTestCase(1).WithArgs(10, []int{3}, []int{3}),
	NewTestCase(1).WithArgs(100, []int{0, 2, 4}, []int{4, 2, 1}),
}

func main() {
	RunTestCases(testCases, func(args ...any) ReturnType {
		return carFleet(args[0].(int), args[1].([]int), args[2].([]int))
	})
}
