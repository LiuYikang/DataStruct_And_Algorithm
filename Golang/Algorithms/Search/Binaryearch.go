package search

func BinarySearch(target int, nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low < high {
		mid := (high-low)/2 + low
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else {
			low = mid
		}
	}
	if nums[low] == target {
		return low
	}
	if nums[high] == target {
		return high
	}
	return -1
}

func BinarySearch2(target int, nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high-low)/2 + low
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid
		} else {
			low = mid
		}
	}
	return -1
}
