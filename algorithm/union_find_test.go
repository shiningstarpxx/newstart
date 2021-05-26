/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  12:42 PM  26/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package algorithm

import (
	"testing"
)

func TestNewUnionFind(t *testing.T) {

	v := NewUnionFind(5)
	if v.Count() != 5 {
		t.Errorf("Unexpect")
	}
}