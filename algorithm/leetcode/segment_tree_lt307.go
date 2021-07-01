/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  5:21 PM  12/6/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package leetcode

type NumArray struct {
	tree []int
	size int
}


func Constructor(nums []int) NumArray {
	n := len(nums)

	v := NumArray{
		tree: make([]int, n * 2),
		size: n,
	}
	for i, num := range nums {
		v.tree[n + i] = num
	}

	for i := n - 1; i > 0; i-- {
		v.tree[i] = v.tree[2 * i] + v.tree[2 * i + 1]
	}
	return v
}


func (this *NumArray) Update(index int, val int)  {
	this.tree[this.size + index] = val
	for i := this.size + index; i > 0; i /= 2 {
		this.tree[i >> 1] = this.tree[i] + this.tree[i ^ 1]
	}
}


/*
 	for this game, [left, right] will be added together
	a little different from basic algorithm [left, right)
 */
func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	l := this.size + left
	r := this.size + right

	for l <= r {
		if l & 1 == 1 {
			sum += this.tree[l]
			l++
		}
		if r & 1 != 1 {
			sum += this.tree[r]
			r--
		}
		l >>= 1
		r >>= 1
	}
	return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
