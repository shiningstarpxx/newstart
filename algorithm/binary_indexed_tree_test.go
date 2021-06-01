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
	r, e := v.GetSum(5)
	a.Equal(12, r)
	a.Nil(e)

	r, e = v.GetSum(100)
	a.NotNil(e)
}

func TestNewBinaryIndexedTree(t *testing.T) {
	freq := []int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}
	v := NewBinaryIndexedTree(freq)
	a := assert.New(t)

	a.Equal(len(freq) + 1, len(v.tree))
}

func TestBinaryIndexedTree_UpdateBITree(t *testing.T) {
	freq := []int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}
	v := NewBinaryIndexedTree(freq)
	a := assert.New(t)

	b := v.tree[4]
	v.UpdateBITree(3, 6)
	a.Equal(b + 6, v.tree[4])
}

func TestBinaryIndexedTree_GetRange(t *testing.T) {
	freq := []int{2, 1, 1, 3, 2, 3, 4, 5, 6, 7, 8, 9}
	v := NewBinaryIndexedTree(freq)
	a := assert.New(t)

	r, e := v.GetRange(0, 5)
	a.Equal(12, r)
	a.Nil(e)

	r, e = v.GetRange(4, 4)
	a.Equal(2, r)
	a.Nil(e)

	r, e = v.GetRange(-1, 4)
	a.NotNil(e)
	r, e = v.GetRange(10, -10)
	a.NotNil(e)
	r, e = v.GetRange(0, 100)
	a.NotNil(e)
	r, e = v.GetRange(100, 5)
	a.NotNil(e)

	r, e = v.GetRange(5, 3)
	a.NotNil(e)
}