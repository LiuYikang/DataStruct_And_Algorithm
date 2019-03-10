package main

import (
	"fmt"
)

/*
quick sort implementation one
*/
func Sort1(nums []int) {
	quickSort1(nums, 0, len(nums)-1)
}

func quickSort1(nums []int, low, high int) {
	if low >= high {
		return
	}
	i, j := low, high
	pivot := nums[i]

	for i < j {
		for i < j && nums[i] < pivot {
			i++
		}
		for i < j && nums[j] > pivot {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	quickSort1(nums, low, i-1)
	quickSort1(nums, i+1, high)
}

/*
quick sort implementation two
*/

func Sort2(nums []int) {
	quickSort2(nums, 0, len(nums)-1)
}

func quickSort2(nums []int, low, high int) {
	if low >= high {
		return
	}
	pivot := nums[high]
	storeIdx := low
	i := low
	for i < high {
		if nums[i] < pivot {
			nums[storeIdx], nums[i] = nums[i], nums[storeIdx]
			storeIdx++
		}
		i++
	}
	nums[storeIdx], nums[high] = nums[high], nums[storeIdx]
	quickSort2(nums, low, storeIdx-1)
	quickSort2(nums, storeIdx+1, high)
}

/*
concurrent quick sort implementation
*/
func concurrentSort(nums []int) {
	chanReceive := make(chan int)
	defer close(chanReceive)
	go concurrentQuickSort(nums, 0, len(nums)-1, chanReceive)
	<-chanReceive
}

func concurrentQuickSort(nums []int, low, high int, chanSend chan int) {
	if low >= high {
		chanSend <- 0
		return
	}

	i := low
	pivot := nums[high]
	storeIdx := low
	for ; i < high; i++ {
		if nums[i] < pivot {
			nums[storeIdx], nums[i] = nums[i], nums[storeIdx]
			storeIdx++
		}
	}
	nums[storeIdx], nums[high] = nums[high], nums[storeIdx]

	chanReceive := make(chan int)
	defer close(chanReceive)
	go concurrentQuickSort(nums, low, storeIdx-1, chanReceive)
	go concurrentQuickSort(nums, storeIdx+1, high, chanReceive)
	<-chanReceive
	<-chanReceive
	chanSend <- 0
	return
}

func main() {
	nums1 := []int{4, 3, 5, 6, 2}
	Sort1(nums1)
	fmt.Println(nums1)

	nums2 := []int{4, 3, 5, 6, 2}
	Sort2(nums2)
	fmt.Println(nums2)

	nums3 := []int{4, 3, 5, 6, 2}
	concurrentSort(nums3)
	fmt.Println(nums3)
}
