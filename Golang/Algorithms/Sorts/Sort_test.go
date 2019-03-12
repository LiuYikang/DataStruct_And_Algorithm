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

func TestQuickSort1(t *testing.T) {
	nums := []int{2, 4, 3, 1, 5, 6}
	Sort1(nums)
	if CheckResult(nums) == false {
		t.Fail()
	}
}

func TestQuickSort2(t *testing.T) {
	nums := []int{2, 4, 3, 1, 5, 6}
	Sort2(nums)
	if CheckResult(nums) == false {
		t.Fail()
	}
}

func TestConcurrentQuickSort(t *testing.T) {
	nums := []int{2, 4, 3, 1, 5, 6}
	concurrentSort(nums)
	if CheckResult(nums) == false {
		t.Fail()
	}
}

func TestBubbleSort(t *testing.T) {
	nums := []int{2, 4, 3, 1, 5, 6}
	BubbleSort(nums)
	if CheckResult(nums) == false {
		t.Fail()
	}
}

func TestSelectSort(t *testing.T) {
	nums := []int{2, 4, 3, 1, 5, 6}
	SelectSort(nums)
	if CheckResult(nums) == false {
		t.Fail()
	}
}

func TestInsertSort(t *testing.T) {
	nums := []int{2, 4, 3, 1, 5, 6}
	InsertSort(nums)
	if CheckResult(nums) == false {
		t.Fail()
	}
}
