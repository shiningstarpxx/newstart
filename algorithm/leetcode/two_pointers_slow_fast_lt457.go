/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: two_pointers_slow_fast_lt457.go
 * @Version: 1.0.0
 * @Data: 2021/8/7 11:41 AM
 */
package leetcode

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(index int) int {
		return ((nums[index] + index) % n + n) % n
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}

		slow := i
		fast := next(i)

		for nums[slow] * nums[fast] > 0 && nums[slow] * nums[next(fast)] > 0 {
			if fast == slow {
				if slow != next(slow) {
					return true
				} else {
					break
				}
			}
			slow = next(slow)
			fast = next(next(fast))
		}

		add := i
		for nums[add] * nums[next(add)] > 0 {
			t := add
			add = next(add)
			nums[t] = 0
		}
	}

	return false
}
