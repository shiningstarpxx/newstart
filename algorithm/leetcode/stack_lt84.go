/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  4:31 PM  24/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"fmt"
	"math"
)

func largestRectangleArea(heights []int) int {
	sentinel := make([]int, len(heights) + 2)
	for i, v := range heights {
		sentinel[i + 1] = v
	}

	res := 0
	stack := []int{0}
	for i := 1; i < len(sentinel); i++ {
		for sentinel[i] < sentinel[stack[len(stack) - 1]] {
			height := sentinel[stack[len(stack) - 1]]
			stack = stack[:len(stack) - 1]
			width := i - stack[len(stack) - 1] - 1
			fmt.Printf("%d, %d\n", height, width)
			res = int(math.Max(float64(res), float64(height * width)))
			fmt.Printf("%v\n", sentinel)
		}
		stack = append(stack, i)
	}
	return res
}