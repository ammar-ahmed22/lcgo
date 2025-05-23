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

## Approach
Basically, we want to check if there are duplicates in each row, column and subbox.

To check for duplicates, we can write a helper function that will use a hashmap technique to check for duplicates.

Checking rows and columns is straightforward, the complexity arises with checking the subboxes because we need to figure out how to iterate over them. To keep it simple, we can do 4 nested loops, the box row, the box column, the inner row, the inner column, then we can index the board with `board[3 * box_row + row][3 * box_col + col]`.
Using this, we can create the array of values for the subbox and check for duplicates.

### Complexity
#### Time: `O(n)`
Since we iterate over the rows, columns and subboxes separately, it's all `O(n)`

#### Space: `O(n)`
The only extra space we create is for the rows, columns and subboxes which all have size of `9`. Technically, it's all constant but we'll give it `O(n)`

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
