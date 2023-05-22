package slice_map_bm

import (
	crand "crypto/rand"
	"fmt"
	mrand "math/rand"
	"testing"
)

const (
	strLenMin = 5
	strLenMax = 50

	sliceSeekLimit = 50000
)

var (
	testSizes []int
)

func BenchmarkSliceDynFill(b *testing.B) {
	for _, size := range testSizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			items := sliceFillDyn(size)
			_ = sliceUse(items)
		})
	}
}

func BenchmarkSliceFixFill(b *testing.B) {
	for _, size := range testSizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			items := sliceFillFix(size)
			_ = sliceUse(items)
		})
	}
}

func BenchmarkSliceDynSeek(b *testing.B) {
	for _, size := range testSizes {
		if size > sliceSeekLimit {
			continue
		}
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			b.StopTimer()
			items := sliceFillDyn(size)
			b.StartTimer()
			for s := 0; s < size; s++ {
				_ = sliceSeek(items, allStrings[s])
			}
		})
	}
}

func BenchmarkSliceFixSeek(b *testing.B) {
	for _, size := range testSizes {
		if size > sliceSeekLimit {
			continue
		}
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			b.StopTimer()
			items := sliceFillFix(size)
			b.StartTimer()
			for s := 0; s < size; s++ {
				_ = sliceSeek(items, allStrings[s])
			}
		})
	}
}

func BenchmarkMapDynFill(b *testing.B) {
	for _, size := range testSizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			items := mapFillDyn(size)
			_ = mapUse(items)
		})
	}
}

func BenchmarkMapFixFill(b *testing.B) {
	for _, size := range testSizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			items := mapFillFix(size)
			_ = mapUse(items)
		})
	}
}

func BenchmarkMapDynSeek(b *testing.B) {
	for _, size := range testSizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			b.StopTimer()
			items := mapFillDyn(size)
			b.StartTimer()
			for s := 0; s < size; s++ {
				_, _ = mapSeek(items, allStrings[s])
			}
		})
	}
}

func BenchmarkMapFixSeek(b *testing.B) {
	for _, size := range testSizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			b.StopTimer()
			items := mapFillFix(size)
			b.StartTimer()
			for s := 0; s < size; s++ {
				_, _ = mapSeek(items, allStrings[s])
			}
		})
	}
}

func init() {
	r := mrand.New(mrand.NewSource(99))
	for i := 0; i < iMax; i++ {
		allStrings[i] = randomText(strLenMin + r.Intn(strLenMax-strLenMin))
	}
	for n := nMin; n <= nMax; n++ {
		testSizes = append(testSizes, 1<<n)
	}
	fmt.Printf("sizes: %v\n", testSizes)
}

func randomText(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz _-,.!?"
	var bytes = make([]byte, n)
	crand.Read(bytes) //nolint:errcheck // not probable
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
