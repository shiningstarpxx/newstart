/*
* @Author Pei, Xingxin
* @Description //TODO $
* @Date $ $
* @Param $
 */
package algorithm

// An IntHeap is a min-heap of ints.
// 迫于没有模板，但是平常的算法解释说明都是

type MiniHeap []int

func (h MiniHeap) Len() int           { return len(h) }
func (h MiniHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MiniHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MiniHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MiniHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

