package bitmap_test

import (
	"fmt"
	"github.com/zgpwr/bitmap"
	"math/rand"
	"testing"
	"time"
)

func TestBitMap(t *testing.T) {
	bm := bitmap.NewBitMap(35)
	bm.SetBit(5)
	bm.SetBit(2)
	fmt.Println(bm.GetBit(3))
	fmt.Println(bm.GetBit(5))
	bm.ClearBit(5)
	fmt.Println(bm.GetBit(5))

	bm1 := bitmap.NewBitMap(10)
	bits := []int{1, 3, 4, 5, 8, 9, 10}
	for _, v := range bits {
		bm1.SetBit(uint32(v))
	}
	fmt.Printf("bm1: %s, bitCount:%d\n", bm1.ToString(), bm1.BitCount())
	fmt.Printf("bm1: data:%v\n", bm1.GetData())

	bm2 := bitmap.NewBitMap(10)
	bits = []int{1, 3, 4, 6, 8, 9, 10}
	for _, v := range bits {
		bm2.SetBit(uint32(v))
	}
	fmt.Printf("bm2: %s, bitCount:%d\n", bm2.ToString(), bm2.BitCount())
	fmt.Printf("bm2: data:%v\n", bm2.GetData())

	bm3 := bm1.And(bm2)
	fmt.Printf("b1 b2 交集: %s, bitCount:%d\n", bm3.ToString(), bm3.BitCount())
	fmt.Printf("bm3: data:%v\n", bm3.GetData())

	bm3 = bm1.Or(bm2)
	fmt.Printf("b1 b2 并集: %s, bitCount:%d\n", bm3.ToString(), bm3.BitCount())
	fmt.Printf("bm3: data:%v\n", bm3.GetData())

	bm3 = bm1.AndNot(bm2)
	fmt.Printf("b1 b2 差集: %s, bitCount:%d\n", bm3.ToString(), bm3.BitCount())
	fmt.Printf("bm3: data:%v\n", bm3.GetData())
}

func TestLargeBitMap(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	bm1 := bitmap.NewBitMap(100000000)
	for i := 0; i < 100000000; i++ {
		num := rand.Int31n(100000000)
		bm1.SetBit(uint32(num))
	}
	fmt.Printf("bm1: bitcount: %d\n", bm1.BitCount())
}
