package testutils

import (
	"fmt"
	"strings"
)

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

func ListFromNums(nums ...int) *ListNode {
	arr := []int{}
	for _, num := range nums {
		arr = append(arr, num)
	}
	return ListFromSlice(arr)
}
