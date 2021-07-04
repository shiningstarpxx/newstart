/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  3:51 PM  4/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func findErrorNums(nums []int) []int {
	return hashCalc(nums)
}

func hashCalc(nums []int) []int {
	cache := map[int]int{}
	for _, v := range nums {
		cache[v]++
	}
	f, s := 0, 0
	for i := 1; i <= len(nums); i++ {
		if c, ok := cache[i]; ok {
			if c == 2 {
				f = i
			}
		} else {
			s = i
		}
		if f != 0 && s != 0 {
			return []int{f, s}
		}
	}
	return []int{f, s}
}

func bitCalc(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}

	for i := 1; i <= len(nums); i++ {
		xor ^= i
	}

	// now we have f ^ s
	lowbit := xor & (-xor)
	one, other := 0, 0
	for _, v := range nums {
		if v & lowbit != 0 {
			one ^= v
		} else {
			other ^= v
		}
	}
	for i := 1; i <= len(nums); i++ {
		if i & lowbit != 0 {
			one ^= i
		} else {
			other ^= i
		}
	}

	for _, v := range nums {
		if v == one {
			return []int{one, other}
		}
	}
	return []int{other, one}
}