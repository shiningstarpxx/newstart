/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  12:42 PM  26/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	a := assert.New(t)
	v := NewUnionFind(5)
	a.Equal(5, v.Count())

	// {0, 1}, 2, 3, 4
	v.Union(0, 1)
	a.Equal(4, v.Count())
}

func TestUnionFind_IsConnected(t *testing.T) {
	v := NewUnionFind(5)

	// {0, 1}, 2, 3, 4
	v.Union(0, 1)

	// {0, 1}, 2, {3, 4}
	v.Union(3, 4)

	// {0, 1, 2}, {3, 4}
	v.Union(2, 1)

	a := assert.New(t)
	a.Equal(true, v.IsConnected(0, 1))
	a.Equal(true, v.IsConnected(0, 2))
	a.Equal(true, v.IsConnected(1, 2))
	a.True(v.IsConnected(3, 4))

	a.False(v.IsConnected(0, 3))
	a.False(v.IsConnected(0, 4))
	a.False(v.IsConnected(1, 3))
	a.False(v.IsConnected(1, 4))
	a.False(v.IsConnected(2, 3))
	a.False(v.IsConnected(2, 4))
}