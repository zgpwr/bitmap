package bitmap

import (
	"strconv"
)

type BitMap struct {
	data     []byte
	bitSize  uint32
	byteSize int
}

func NewBitMap(size uint32) *BitMap {
	if size%8 > 0 {
		size += 8 - size%8
	}
	byteSize := size >> 3
	return &BitMap{
		data:     make([]byte, byteSize),
		bitSize:  size,
		byteSize: int(byteSize),
	}
}

func (bm *BitMap) SetBit(offset uint32) {
	if offset > bm.bitSize {
		return
	}
	offset--
	index, pos := offset>>3, offset%8
	bm.data[index] |= 0x01 << pos
}

func (bm *BitMap) GetBit(offset uint32) uint8 {
	if offset > bm.bitSize {
		return 0
	}
	offset--
	index, pos := offset>>3, offset%8
	return bm.data[index] >> pos & 0x01
}

func (bm *BitMap) ClearBit(offset uint32) {
	if offset > bm.bitSize {
		return
	}
	offset--
	index, pos := offset>>3, offset%8
	bm.data[index] &^= 0x01 << pos
}

func (bm *BitMap) Clone() *BitMap {
	bmCopy := NewBitMap(bm.bitSize)
	copy(bmCopy.data, bm.data)
	return bmCopy
}

func (bm *BitMap) And(other *BitMap) *BitMap {
	size := min(bm.bitSize, other.bitSize)
	res := NewBitMap(size)
	for i := 0; i < res.byteSize; i++ {
		res.data[i] = bm.data[i] & other.data[i]
	}

	return res
}

func (bm *BitMap) Or(other *BitMap) *BitMap {
	var res, s *BitMap
	if bm.byteSize > other.byteSize {
		res = bm.Clone()
		s = other
	} else {
		res = other.Clone()
		s = bm
	}
	for i := 0; i < res.byteSize; i++ {
		res.data[i] |= s.data[i]
	}

	return res
}

func (bm *BitMap) AndNot(other *BitMap) *BitMap {
	bmCopy := bm.Clone()
	for i := 0; i < bmCopy.byteSize; i++ {
		bmCopy.data[i] &^= other.data[i]
	}

	return bmCopy
}

func (bm *BitMap) ToString() string {
	res := ""
	for i := bm.bitSize; i > 0; i-- {
		res += strconv.Itoa(int(bm.GetBit(i)))
	}
	return res
}

func (bm *BitMap) GetBitSize() uint32 {
	return bm.bitSize
}

func (bm *BitMap) GetByteSize() int {
	return bm.byteSize
}

func (bm *BitMap) GetData() []byte {
	return bm.data
}

func min(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

