/*
* @Author Pei, Xingxin
* @Description //TODO $
* @Date $ $
* @Param $
 */
package example

import (
	"container/heap"
	al "mic_test/algorithm"
)

type MedianFinder struct {
	low al.MaxHeap
	high al.MiniHeap
}

// Constructor295 /** initialize your data structure here. */
func Constructor295() MedianFinder {
	v := MedianFinder{
		make([]int, 0),
		make([]int, 0),
	}
	return v
}

func (this *MedianFinder) AddNum(num int)  {
	if this.low.Len() == this.high.Len() {
		heap.Push(&this.low, num)
	} else {
		heap.Push(&this.high, num)
	}
	// fmt.Printf("%v, %v\n", this.low, this.high)

	for this.high.Len() > 0 && this.low[0] > this.high[0] {
		i, j := heap.Pop(&this.low), heap.Pop(&this.high)
		// fmt.Printf("%d, %d\n", i, j)
		heap.Push(&this.low, j)
		heap.Push(&this.high, i)
	}
	// fmt.Printf("adjust %v, %v\n", this.low, this.high)
}


func (this *MedianFinder) FindMedian() float64 {
	if this.low.Len() > this.high.Len() {
		return float64(this.low[0])
	} else {
		return float64(this.low[0] + this.high[0]) / 2
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */