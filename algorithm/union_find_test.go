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

