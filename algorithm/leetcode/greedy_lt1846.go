/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  2:24 PM  15/7/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"math"
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	data := append(sort.IntSlice{}, arr...)
	data.Sort()

	data[0] = 1
	for i := 1; i < len(data); i++ {
		if data[i] > data[i - 1] + 1 {
			data[i] = data[i - 1] + 1
		}
	}

	for i := len(data) - 2; i >= 0; i-- {
		if data[i] > data[i + 1] + 1 {
			data[i] = data[i + 1] + 1
		}
	}

	m := 0
	for _, v := range data {
		m = int(math.Max(float64(m), float64(v)))
	}
	return m
}
