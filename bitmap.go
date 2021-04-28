package bitmap

import (
	"strconv"
)

type BitMap struct {
	data      []byte
	bitSize   uint32
	maxOffset uint32
}

func NewBitMap(size uint32) *BitMap {
	if size%8 > 0 {
		size += 8 - size%8
	}
	return &BitMap{
		bitSize: size,
		data:    make([]byte, size>>3),
	}
}

func (bm *BitMap) SetBit(offset uint32) {
	index, pos := offset>>3, offset%8
	bm.data[index] |= 0x01 << pos
	if bm.maxOffset < offset {
		bm.maxOffset = offset
	}
}

func (bm *BitMap) GetBit(offset uint32) uint8 {
	index, pos := offset>>3, offset%8
	return bm.data[index] >> pos & 0x01
}

func (bm *BitMap) GetData() string {
	res := ""
	for i := bm.maxOffset; i > 0; i-- {
		res += strconv.Itoa(int(bm.GetBit(i)))
	}
	return res
}

func (bm *BitMap) GetSize() uint32 {
	return bm.bitSize
}
