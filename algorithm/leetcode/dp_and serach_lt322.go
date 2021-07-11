/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:06 AM  11/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	return coinChangeDP(coins, amount)
}

func coinChangeDP(coins []int, amount int) int {
	cache := make([]int, amount + 1)

	for i := range cache {
		cache[i] = math.MaxInt32
	}
	cache[0] = 0

	for _, c := range coins {
		for tip := c; tip <= amount; tip++ {
			cache[tip] = int(math.Min(float64(cache[tip]), float64(cache[tip - c] + 1)))
		}
	}

	if cache[amount] != math.MaxInt32 {
		return cache[amount]
	}
	return -1
}

// @todo(michael): time limited exceed
func coinChangeDFS(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	fmt.Printf("%v\n", coins)

	res := cddfs(coins, 0, amount, 0, math.MaxInt32)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func cddfs(coins []int, index int, amount, use, ans int) int {
	if index == len(coins) - 1 || amount == 0 {
		if amount % coins[index] == 0 {
			return amount / coins[index] + use
		}
		return math.MaxInt32
	}

	for k := amount / coins[index]; k >= 0 && k + use < ans; k-- {
		n := cddfs(coins, index + 1, amount - k * coins[index], use + k, ans)
		ans = int(math.Min(float64(ans), float64(n)))
	}
	return ans
}