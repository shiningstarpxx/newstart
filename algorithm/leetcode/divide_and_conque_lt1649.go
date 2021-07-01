/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  3:44 PM  17/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

func createSortedArray(instructions []int) int {
	sorted := make([]int, len(instructions))
	copy(sorted, instructions)
	smallerCount := make([]int, len(instructions))
	createSortedArrayHelper(instructions, sorted, smallerCount, 0, len(instructions) - 1)
	res := 0
	count := make(map[int]int)
	for i, v := range instructions {
		res += min(smallerCount[i], i - smallerCount[i] - count[v])
		res %= 1e9+7
		count[v]++
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func createSortedArrayHelper(nums, sorted, count []int, l, r int) {
	if l == r {
		return
	}
	mid := l + (r - l) / 2
	createSortedArrayHelper(nums, sorted, count, l, mid)  // sorted[l...mid] are ordered
	createSortedArrayHelper(nums, sorted, count, mid + 1, r)  // sorted[mid+1...r] are ordered

	for i := mid + 1; i <= r; i++ {
		index := LowerBound1649(sorted, l, mid, nums[i])
		// fmt.Printf("%v\n", sorted[l:mid+1])
		// fmt.Printf("%d, %d, %d, %d, %d, %d\n", i, l, mid, r, nums[i], index)
		count[i] += index - l
	}

	merge(sorted, l, mid, r)
	// fmt.Printf("%v\n", count)
}

func merge(sorted []int, l, mid, r int)  {
	t := make([]int, r - l + 1)
	index, i, j :=0, l, mid + 1
	for i <= mid || j <= r {
		if i > mid {
			t[index] = sorted[j]
			j++
		} else if j > r {
			t[index] =  sorted[i]
			i++
		} else if sorted[i] <= sorted[j] {
			t[index] = sorted[i]
			i++
		} else {
			t[index] = sorted[j]
			j++
		}
		index++
	}
	for i := 0; i < len(t); i++ {
		sorted[l + i] = t[i]
	}
}

func LowerBound1649(arr []int, left, right int, target int) int {
	l, r := left, right
	for ; l <= r; {
		mid := l + (r-l)/2
		if arr[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}