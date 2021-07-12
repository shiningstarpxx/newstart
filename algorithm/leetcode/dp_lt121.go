/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  11:01 AM  12/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import "math"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minPrice := prices[0]  // till pre-day min price
	maxPrice := 0  // till today, most value

	for i := 1; i < len(prices); i++ {
		maxPrice = int(math.Max(float64(prices[i] - minPrice), float64(maxPrice)))
		minPrice = int(math.Min(float64(prices[i]), float64(minPrice)))
	}
	return maxPrice
}