/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: slice.go
 * @Version: 1.0.0
 * @Data: 2022/3/11 10:58 PM
 */
package leveldb

import (
	"errors"
	"strings"
)

type LSlice struct {
	size int
	data string
}

func NewEmptySlice() *LSlice {
	return &LSlice{
		size: 0,
		data: "",
	}
}

func NewSlice(s string, n int) *LSlice {
	return &LSlice{
		size: n,
		data: s,
	}
}

func NewSliceWithStr(s string) *LSlice {
	return &LSlice{
		size: len(s),
		data: s,
	}
}

func (s *LSlice) GetData() string {
	return s.data
}

func (s *LSlice) GetSize() int {
	return s.size
}

func (s *LSlice) Empty() bool {
	return s.size == 0
}

func (s *LSlice) Clear() {
	s.size = 0
	s.data = ""
}

func (s *LSlice) RemovePrefix(n int) {
	if n > s.size {
		panic(errors.New("remove more size than had"))
	}
	s.data = s.data[n:]
	s.size -= n
}

func (s *LSlice) StartWith(p *LSlice) bool {
	return s.size >= p.GetSize() && strings.Compare(string(s.data[:p.size]), p.GetData()) == 0
}

func (s *LSlice) Compare(p *LSlice) int {
	minLen := 0
	if s.size <= p.GetSize() {
		minLen = s.size
	} else {
		minLen = p.GetSize()
	}

	r := strings.Compare(s.data[:minLen], p.GetData()[:minLen])
	if r == 0 {
		if s.size < p.GetSize() {
			return -1
		} else if s.size > p.GetSize() {
			return 1
		}
	}
	return r
}