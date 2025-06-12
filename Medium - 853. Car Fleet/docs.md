# 853. Car Fleet

## Problem

There are `n` cars at given miles away from the starting mile 0, traveling to reach the mile `target`.

You are given two integer array `position` and `speed`, both of length `n`, where `position[i]` is the starting mile of the `ith` car and `speed[i]` is the speed of the `ith` car in miles per hour.

A car cannot pass another car, but it can catch up and then travel next to it at the speed of the slower car.

A **car fleet** is a car or cars driving next to each other. The speed of the car fleet is the **minimum** speed of any car in the fleet.

If a car catches up to a car fleet at the mile `target`, it will still be considered as part of the car fleet.

Return the number of car fleets that will arrive at the destination.

**Example 1:**

**Input:** target = 12, position = \[10,8,0,5,3\], speed = \[2,4,1,1,3\]

**Output:** 3

**Explanation:**

- The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12. The fleet forms at `target`.
- The car starting at 0 (speed 1) does not catch up to any other car, so it is a fleet by itself.
- The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches `target`.

**Example 2:**

**Input:** target = 10, position = \[3\], speed = \[3\]

**Output:** 1

**Explanation:**

There is only one car, hence there is only one fleet.

**Example 3:**

**Input:** target = 100, position = \[0,2,4\], speed = \[4,2,1\]

**Output:** 1

**Explanation:**

- The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting each other at 4. The car starting at 4 (speed 1) travels to 5.
- Then, the fleet at 4 (speed 2) and the car at position 5 (speed 1) become one fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches `target`.

**Constraints:**

- `n == position.length == speed.length`
- `1 <= n <= 105`
- `0 < target <= 106`
- `0 <= position[i] < target`
- All the values of `position` are **unique**.
- `0 < speed[i] <= 106`

## Approach
To solve this problem, we can use a stack-based approach. 

As mentioned in the problem description, the general idea is that once a car hits another car, it reduces it's speed to match that cars speed. Therefore, we want to process the cars in sorted order of their positions, so, the first step would be to sort the arrays by position. To make this easier, we'll combined the position and speed of a car and then sort the combined array by position.

After this, we want to check the time it will take a given car to reach the target, if the car before it will reach the target faster than it, then that means the previous car will eventually match up with the car ahead and create a fleet.

Therefore, to find the solution, we can traverse the combined, sorted array in reverse order alongside a stack. On each iteration, we calculate the `timeToTarget` for the given car, if the `timeToTarget`for that car is less than or equal to the top of the stack, we don't add it to the stack because it will match up with the car ahead. Otherwise, we add to the stack and continue. 

By the end, our stack will be filled with the time's it will take the limiting cars to reach the target so our answer will simply be the length of the stack.

## Complexity
### Time: `O(nlogn)`
We start off by sorting the input arrays which is `O(nlogn)`. After that we iterate over the combined array once which is `O(n)` giving us `O(n) + O(nlogn) = O(2nlogn) = O(nlogn)`

### Space: `O(n)`
We create the stack (`O(n)`) and the combined array `O(n)` giving us `O(n) + O(n) = O(2n) = O(n)`

## Solution

```go
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

```
