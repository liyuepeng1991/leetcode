package twosum

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 11, 7, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
