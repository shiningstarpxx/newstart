/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: slice.go
 * @Version: 1.0.0
 * @Data: 2022/3/11 10:58 PM
 */
package leveldb

type Slice struct {
	size int
	data string
}

func NewEmptySlice() *Slice {
	return &Slice{
		size: 0,
		data: "",
	}
}

func NewSlice(s string, n int) *Slice {
	return &Slice{
		size: n,
		data: s,
	}
}

func NewSliceWithStr(s string) *Slice {
	return &Slice{
		size: len(s),
		data: s,
	}
}

func (s *Slice) GetData() string {
	return s.data
}