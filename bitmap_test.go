package bitmap_test

import (
	"fmt"
	"github.com/zgpwr/bitmap"
	"testing"
)

func TestBitMap(t *testing.T) {
	bm := bitmap.NewBitMap(35)
	fmt.Println(bm.GetSize())
	bm.SetBit(5)
	bm.SetBit(2)
	fmt.Println(bm.GetBit(3))
	fmt.Println(bm.GetBit(5))
	fmt.Println(bm.GetData())
}
