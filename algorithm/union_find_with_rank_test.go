/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  10:08 PM  29/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUnionFindWithRank(t *testing.T) {
	a := assert.New(t)
	v := NewUnionFindWithRank(5)
	a.Equal(5, v.Count())

	// {0, 1}, 2, 3, 4
	v.Union(0, 1)
	a.Equal(4, v.Count())
}

func TestUnionFindWithRank_IsConnected(t *testing.T) {
	v := NewUnionFindWithRank(5)

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

func TestUnionFindWithRank_Find(t *testing.T) {
	v := NewUnionFindWithRank(5)
	a := assert.New(t)

	// {0->1}, 2, 3, 4
	v.Union(0, 1)

	// {0, 1}, 2, {3->4}
	v.Union(3, 4)

	// {0->1->2}, {3->4}
	v.Union(2, 1)

	a.Equal(1, v.Find(1))
	a.Equal(1, v.Find(0))
	a.Equal(1, v.Find(2))
	a.Equal(4, v.Find(4))
	a.Equal(4, v.Find(3))

	// {0->1->2->3->4}
	// In fact, after compaction
	// {0, 1, 2, 3} -> {4}
	v.Union(0, 3)
	a.Equal(4, v.Find(0))
}

func TestUnionFindWithRank_Union(t *testing.T) {
	v := NewUnionFindWithRank(5)
	a := assert.New(t)

	// {0, 1}, 2, 3, 4
	v.Union(0, 1)
	a.Equal(4, v.Count())
	// TODO(michael), Need find a way to test lacal value
	a.Equal(1, v.Find(0))
	a.Equal(1, v.Find(1))

	// {0, 1}, 2, {3, 4}
	v.Union(3, 4)
	a.Equal(3, v.Count())
	a.Equal(4, v.Find(3))
	a.Equal(4, v.Find(4))

	// {0, 1, 2}, {3, 4}
	v.Union(2, 1)
	a.Equal(2, v.Count())
	a.Equal(1, v.Find(0))
	a.Equal(1, v.Find(1))
	a.Equal(1, v.Find(2))

	// Repeat Union nothing change
	v.Union(2, 1)
	a.Equal(2, v.Count())

	// {0, 1, 2, 3, 4}
	v.Union(0, 4)
	a.Equal(1, v.Count())
	a.Equal(4, v.Find(0))
	a.Equal(4, v.Find(1))
	a.Equal(4, v.Find(2))
	a.Equal(4, v.Find(3))
	a.Equal(4, v.Find(4))
}