/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  4:37 PM  30/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestConstructor295(t *testing.T) {
	tests := []struct {
		name string
		want MedianFinder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor295(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor295() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
[[],[1],[2],[],[3],[]]
 */
func TestMedianFinder(t *testing.T) {
	a := assert.New(t)

	v := Constructor295()
	v.AddNum(1)
	v.AddNum(2)
	a.Equal(float64(1 + 2)/2, v.FindMedian())
	v.AddNum(3)

	a.Equal(float64(2), v.FindMedian())
}

/*
["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
[[],[-1],[],[-2],[],[-3],[],[-4],[],[-5],[]]
 */
func TestMedianFinder_EdgeCase(t *testing.T) {
	a := assert.New(t)

	v := Constructor295()
	v.AddNum(-1)
	a.Equal(float64(-1), v.FindMedian())
	v.AddNum(-2)
	a.Equal(float64(-1 - 2)/2, v.FindMedian())
	v.AddNum(-3)
	a.Equal(float64(-2), v.FindMedian())
	v.AddNum(-4)
	a.Equal(float64(-3 - 2)/2, v.FindMedian())
	v.AddNum(-5)
	a.Equal(float64(-3), v.FindMedian())
}