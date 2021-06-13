/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  5:04 PM  30/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package main

import (
	"fmt"
	"strings"
)

func main()  {
	if strings.Contains("a==b", "==") {
		r := strings.Split("a==b", "==")
		for _, v := range r {
			fmt.Printf("%v\n", v)
			fmt.Printf("%v\n", v[0]-'a')
		}
	}

	nums := []int{1, 2, 3, 4, 5}
	l := nums[:3]
	r := nums[3:]
	fmt.Printf("%v\n", l)
	fmt.Printf("%v\n", r)
}