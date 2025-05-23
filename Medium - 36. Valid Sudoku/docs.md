# 36. Valid Sudoku

## Problem

Determine if a `9 x 9` Sudoku board is valid. Only the filled cells need to be validated **according to the following rules**:

1. Each row must contain the digits `1-9` without repetition.
2. Each column must contain the digits `1-9` without repetition.
3. Each of the nine `3 x 3` sub-boxes of the grid must contain the digits `1-9` without repetition.

**Note:**

- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

**Example 1:**

![](https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)

```
Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true

```

**Example 2:**

```
Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

```

**Constraints:**

- `board.length == 9`
- `board[i].length == 9`
- `board[i][j]` is a digit `1-9` or `'.'`.

## Solution Notes
- We can start by checking each row and column
- Essentially we are checking for any duplicates in the row or column ignoring the '.'
    + We can use a hash map to check for duplicates by adding values already seen to the map
- Iterate over the rows, if any have duplicates, return `false`
- Iterate over the columns, if any have duplicates, return `false`
- Now we need to iterate over the sub-boxes
    + There are 9 sub-boxes
    + Iterate over each subbox with two nested loops from `0 to 3` each
    + Inside that do two more nested loops to iterate over the values in each sub-box
        + To get the index on the main board -> `board[3 * brow + row][3 * bcol + col]`
        + Where the `brow` and `bcol` are the box row and col and the `row` and `col` are the inner ones
        + In total, we have 4 nested loops each going from `0 to 3`

### Complexity
#### Time: `O(n)`
- We iterate over the rows and then in the check for duplicates iterate over all the values -> `O(n)`
- We iterater over the columns and the in the check for duplicates iterate over all the avlues -> `O(n)`
- We iterate over the 9 subboxes and then in the check for duplicates iterate over all the values -> `O(n)`

#### Space: `O(n)`
- We create rows and column slices when doing the checks, each of which have a length of `9`
- Would technically be constant but we'll give it `n` since even the input is technically constant size

## Solution

```go
func hasDuplicate(values []byte) bool {
	seen := make(map[byte]bool)
	for _, val := range values {
		if val == '.' {
			continue
		}
		if _, exists := seen[val]; exists {
			return true
		}
		seen[val] = true
	}
	return false
}
func isValidSudoku(board [][]byte) bool {
	n := len(board)
	// Checking rows
	for _, row := range board {
		if hasDuplicate(row) {
			return false
		}
	}

	// Checking columns
	for i := range n {
		col := make([]byte, n)
		for j := range n {
			col[j] = board[j][i]
		}
		if hasDuplicate(col) {
			return false
		}
	}

	// Checking sub-boxes
	for brow := range 3 {
		for bcol := range 3 {
			var values []byte
			for row := range 3 {
				for col := range 3 {
					values = append(values, board[3 * brow + row][3 * bcol + col])
				}
			}
			if hasDuplicate(values) {
				return false
			}
		}
	}

	return true
}

```
