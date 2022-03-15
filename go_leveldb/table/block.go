/**
 * @Author: pei, xingxin
 * @mail: xingxinpei@gmail.com
 * @Description:
 * @File: block.go
 * @Version: 1.0.0
 * @Data: 2022/3/15 3:23 PM
 */
package table

import (
	"my_module_name/include/leveldb"
	"my_module_name/util"
	"unsafe"
)

type BlockContent struct {
	data leveldb.LSlice         // Actual contents of data
	cachable bool        		// True iff data can be cached
	heapAllocated bool  		// True iff caller should delete[] data.data()
}

type Block struct {
	data string
	size uint32
	restartOffset uint32
	owned bool
}

func (b *Block) GetSize() uint32 {
	return b.size
}

func (b *Block) NumRestarts() uint32 {
	if b.GetSize() < 4 {
		panic("block size is too small")
	}
	return util.DecodeFixed32(b.data[b.size - uint32(unsafe.Sizeof(uint32(0))):])
}