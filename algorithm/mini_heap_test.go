/*
* @Author Pei, Xingxin
* @Description //TODO $
* @Date $ $
* @Param $
 */
package algorithm

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMiniHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    MiniHeap
		want int
	}{
		// TODO: Add test cases.
		{
			name : "empty heap",
			h : MiniHeap{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiniHeap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    MiniHeap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiniHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    MiniHeap
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiniHeap_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		h    MiniHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestMiniHeap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    MiniHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestMiniHeap(t *testing.T) {
	h := &MiniHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)

	a := assert.New(t)
	a.Equal(1, (*h)[0])
	a.Equal(4, h.Len())
	a.Equal(1, heap.Pop(h).(int))
	a.Equal(2, (*h)[0])
	a.Equal(2, heap.Pop(h).(int))
	a.Equal(3, heap.Pop(h).(int))
	a.Equal(5, heap.Pop(h).(int))
	a.Equal(0, h.Len())
}
