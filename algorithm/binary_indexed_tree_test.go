/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:58 AM  1/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryIndexedTree_GetSum(t *testing.T) {
	freq := []int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}
	v := NewBinaryIndexedTree(freq)

	a := assert.New(t)
	a.Equal(12, v.GetSum(5))
}