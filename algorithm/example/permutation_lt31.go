/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:35 PM  23/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package example

func nextPermutation(nums []int)  {
	firstNotDecrease := len(nums) - 2
	for firstNotDecrease >= 0 && nums[firstNotDecrease] >= nums[firstNotDecrease + 1] {
		firstNotDecrease--
	}

	if firstNotDecrease >= 0 {
		firstBigger := len(nums) - 1
		for firstBigger > 0 && nums[firstBigger] <= nums[firstNotDecrease] {
			firstBigger--
		}
		nums[firstNotDecrease], nums[firstBigger] = nums[firstBigger], nums[firstNotDecrease]
	}
	i := firstNotDecrease + 1
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}