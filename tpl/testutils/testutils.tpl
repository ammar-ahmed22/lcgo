package testutils

import (
	"fmt"
	"reflect"
    "strings"

	"github.com/fatih/color"
	"github.com/samber/lo"
)

type TestCase[R any] struct {
	Args      []any
	Returns   R
	CompareFn func(R, R) bool
	Name      string
}

func NewTestCase[R any](returns R) *TestCase[R] {
	return &TestCase[R]{Name: "", Returns: returns, Args: []any{}, CompareFn: func(r R, a R) bool {
		return reflect.DeepEqual(r, a)
	}}
}

func RunTestCases[R any](testCases []*TestCase[R], f func(...any) R) error {
	for i, testCase := range testCases {
		var name string
		if testCase.Name != "" {
			name = testCase.Name
		} else {
			name = fmt.Sprintf("Test %d", i+1)
		}
		testCase.Run(name, f)
	}
	return nil
}

func Test[R comparable](name string, expected, received R) {
	if expected == received {
		color.Green("PASSED - %s", name)
	} else {
		color.Red("FAILED - %s", name)
		color.Red("  Expected: %v", expected)
		color.Red("  Received: %v", received)
	}
}

func TestCompare[R any](name string, expected, received R, compareFn func(R, R) bool) {
	if compareFn(expected, received) {
		color.Green("PASSED - %s", name)
	} else {
		color.Red("FAILED - %s", name)
		color.Red("  Expected: %v", expected)
		color.Red("  Received: %v", received)
	}
}

func (t *TestCase[R]) Run(name string, f func(...any) R) {
	TestCompare(name, t.Returns, f(t.Args...), t.CompareFn)
}

func (t *TestCase[R]) WithArgs(args ...any) *TestCase[R] {
	t.Args = args
	return t
}

func (t *TestCase[R]) WithCompareFn(compareFn func(R, R) bool) *TestCase[R] {
	t.CompareFn = compareFn
	return t
}

func (t *TestCase[R]) WithName(name string) *TestCase[R] {
	t.Name = name
	return t
}

func SliceEqualUnordered[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	countA := lo.CountValues(a)
	countB := lo.CountValues(b)
	return reflect.DeepEqual(countA, countB)
}

type ListNode struct {
  Val int
  Next *ListNode
}

func (l *ListNode) String() string {
	var builder strings.Builder
	curr := l
	for curr != nil {
		builder.WriteString(fmt.Sprintf("%d -> ", curr.Val))
		curr = curr.Next
	}
	builder.WriteString("nil")
	return builder.String()
}

func ListFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return &ListNode{}
	}
	list := ListNode{
		Val: nums[0],
	}
	curr := &list
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{
			Val: nums[i],
		}
		curr = curr.Next
	}
	return &list
}
