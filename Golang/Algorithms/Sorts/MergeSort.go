package sort

func MergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	key := n / 2
	left := MergeSort(nums[:key])
	right := MergeSort(nums[key:])
	return merge(left, right)
}

func merge(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	if i < len(nums1) {
		res = append(res, nums1[i:]...)
	}

	if j < len(nums2) {
		res = append(res, nums2[j:]...)
	}
	return res
}
