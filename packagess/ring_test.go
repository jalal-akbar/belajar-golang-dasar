package packagess

import (
	"container/ring"
	"fmt"
	"strconv"
	"testing"
)

func TestRing(t *testing.T) {
	var rings *ring.Ring = ring.New(5)

	for i := 0; i < rings.Len(); i++ {
		rings.Value = "Value Ke " + strconv.FormatInt(int64(i), 10)
		rings = rings.Next()
	}

	rings.Do(func(a any) {
		fmt.Println(a)
	})
}
