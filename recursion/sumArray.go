package recursion

// SumArray 计算数组之和
func SumArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return nums[0] + SumArray(nums[1:])
}

// ArrayLength 计算数组长度
func ArrayLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	return 1 + ArrayLength(nums[1:])
}

// 快排
func QuickSort(nums []int) []int {
	l := len(nums)
	if l < 2 {
		return nums
	}
	pivot := nums[0]
	var small, big []int
	for i := 1; i < l; i++ {
		if nums[i] < pivot {
			small = append(small, nums[i])
		} else {
			big = append(big, nums[i])
		}
	}
	small = QuickSort(small)
	big = QuickSort(big)
	return append(append(small, pivot), big...)
}
