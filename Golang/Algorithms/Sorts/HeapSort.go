package sort

func heapSort(nums []int) {
	length := len(nums)
	first := length/2 - 1
	for i := first; i >= 0; i-- {
		heapify(nums, i, length-1)
	}

	for end := length - 1; end >= 0; end-- {
		nums[end], nums[0] = nums[0], nums[end]
		heapify(nums, 0, end-1)
	}
}

func heapify(nums []int, start, end int) {
	root := start
	for {
		child := 2*root + 1
		if child > end {
			break
		}

		if child+1 <= end && nums[child] < nums[child+1] {
			child++
		}

		if nums[root] < nums[child] {
			nums[root], nums[child] = nums[child], nums[root]
			root = child
		} else {
			break
		}
	}
}
