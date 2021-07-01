/*
   @Copyright: Tencent (1998-2022)
   @Author: Pei, Xingxin
   @Created:  9:07 AM  24/5/2021
   @Email: xingxinpei@gmail, michaelpei@tencent.com
   @Description:
*/
package main

import (
	"errors"
	"fmt"
	"sort"

	"mic_test/algorithm/leetcode"
)

type BinaryIndexedTree struct {
	tree []int
}

func NewBinaryIndexedTreeSize(size int) *BinaryIndexedTree  {
	v := &BinaryIndexedTree{
		tree: make([]int, size + 1),
	}
	return v
}

func (b* BinaryIndexedTree) UpdateBITree(index int, v int) {
	index += 1
	for index < len(b.tree) {
		b.tree[index] += v
		index += index & (-index)
	}
}

func (b* BinaryIndexedTree) GetSum(index int) (int, error){
	if index >= len(b.tree) - 1 || index < 0 {
		return 0, errors.New("out of tree range")
	}
	sum := 0
	index += 1
	for index > 0 {
		sum += b.tree[index]
		index = index - (index & (-index))
	}
	return sum, nil
}

func countSmallerBIT(nums []int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	cache := make(map[int]int, len(nums))
	for i, v := range sorted {
		cache[v] = i + 1
	}

	bit := NewBinaryIndexedTreeSize(len(nums))
	reverseSlice(nums)

	res := make([]int, 0)
	for _, v := range nums {
		sum, _ := bit.GetSum(cache[v] - 1)
		res = append(res, sum)
		bit.UpdateBITree(cache[v], 1)
	}
	reverseSlice(res)
	return res
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	num := []int{1,3,3,3,2,4,2,1,2}
	r := leetcode.CreateSortedArrayST(num)
	fmt.Printf("%v\n", r)
	/*
	t := []int {5, 2, 6, 1}
	r := countSmallerBIT(t)
	fmt.Println(r)

	t1 := []int {-1, -1}
	r = countSmallerBIT(t1)
	fmt.Println(r)
	*/

	/*
	lastAccess := expvar.NewString("lastAccess")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
		lastAccess.Set(time.Now().String())
	})

	http.ListenAndServe(":8080", nil)
	 */
}

// http :8080/debug/vars | grep lastAccess
// 在 expvar.go 中有init，注册了debug/vars，会发布所有Publish的变量
/*
func init() {
        http.HandleFunc("/debug/vars", expvarHandler)
        Publish("cmdline", Func(cmdline))
        Publish("memstats", Func(memstats))
}
 */