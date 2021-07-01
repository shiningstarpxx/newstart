/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  6:04 PM  13/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	sum := make([]int, n + 1)
	for i, v := range nums {
		sum[i + 1] = sum[i] + v
	}
	// fmt.Printf("%v\n", sum)
	return countRangeSumHelper(sum, lower, upper, 0, n)
}

func countRangeSumHelper(sum []int, lower, upper, l, r int) int {
	if l == r {
		return 0
	}

	// [l || r] is our divided and conque way
	// the results are composed by three situtations
	// 1, 2 located in l or r without overlap l->r
	// 3 is the lower is in l, and upper is in r
	mid := l + (r - l) / 2
	res := countRangeSumHelper(sum, lower, upper, l, mid) +
		countRangeSumHelper(sum, lower, upper, mid + 1, r)

	for pivot, rl, rr := l, mid+1, mid+1; pivot <= mid; pivot++ {
		for rl <= r && sum[rl] - sum[pivot] < lower {
			rl++
		}
		for rr <= r && sum[rr] - sum[pivot] <= upper {
			rr++
		}
		res += (rr -rl)
	}

	// todo(michael): Try to express with more elegant way
	// merge two sorted array
	sorted := make([]int, r - l + 1)
	i, ll, rr := 0, l, mid + 1
	for ll <= mid && rr <= r {
		if sum[ll] <= sum[rr] {
			sorted[i] = sum[ll]
			i++
			ll++
		} else {
			sorted[i] = sum[rr]
			i++
			rr++
		}
	}
	for ll <= mid {
		sorted[i] = sum[ll]
		i++
		ll++
	}
	for rr <= r {
		sorted[i] = sum[rr]
		i++
		rr++
	}
	for i := 0; i < len(sorted); i++ {
		sum[i + l] = sorted[i]
	}
	return res
}

// sum[j - i]
// lower <= presum[j] - presum[i] <= upper
// try to solve all dc problems
