package leetcode


func missingNumber(nums []int) int {
	return mathWay(nums)
}

func mathWay(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < n; i++ {
		sum -= nums[i]
	}
	return sum
}
