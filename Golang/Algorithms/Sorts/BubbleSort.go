package sort

func BubbleSort(nums []int) {
	flag := false
	for i := len(nums) - 1; i > 0; i-- {
		flag = true
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
