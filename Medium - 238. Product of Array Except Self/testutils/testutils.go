package testutils

import (
	"fmt"
	"reflect"

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
		success, err := testCase.Run(f)
		if err != nil {
			return err
		}
		var name string
		if testCase.Name != "" {
			name = testCase.Name
		} else {
			name = fmt.Sprintf("Test %d", i+1)
		}

		if success {
			color.Green("PASSED - %s", name)
		} else {
			received := f(testCase.Args...)
			color.Red("FAILED - %s", name)
			color.Red("  Expected: %v", testCase.Returns)
			color.Red("  Received: %v", received)
		}
	}
	return nil
}

func (t *TestCase[R]) Run(f func(...any) R) (bool, error) {
	return t.CompareFn(f(t.Args...), t.Returns), nil
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
