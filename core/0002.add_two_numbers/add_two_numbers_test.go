package addtwonumbers

import (
	"leetcode/util/linklist"
	"testing"
)

func Test_AddTwoNumbers(t *testing.T) {
	a := &linklist.LinkNode{Value: 2}
	a.Add(4)
	a.Add(3)

	b := &linklist.LinkNode{Value: 5}
	b.Add(6)
	b.Add(4)
	c := addtwoNumbers(a, b)
	c.Print()
}
