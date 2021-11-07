package leetcode

func maximumDifference(nums []int) int {
	ans := -1
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > min {
			ans = maxInt2016(ans, nums[i] - min)
		} else {
			min = nums[i]
		}
	}
	return ans
}

func maxInt2016(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bruteForce(nums []int) int {
	n := len(nums)

	ans := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] && nums[j] - nums[i] > ans {
				ans = nums[j] - nums[i]
			}
		}
	}
	return ans
}
