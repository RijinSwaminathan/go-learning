package arrayproblems

func RemoveDuplicates(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i] == nums[j] {
				nums = append(nums, nums[i])
			}
		}
	}
	return nums
}
