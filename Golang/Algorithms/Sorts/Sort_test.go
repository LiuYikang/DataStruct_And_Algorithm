package sort

import (
	"testing"
)

func CheckResult(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func TestHeapSort(t *testing.T) {
	nums := []int{2, 4, 3, 1, 5, 6}
	heapSort(nums)
	if CheckResult(nums) == false {
		t.Fail()
	}
}
