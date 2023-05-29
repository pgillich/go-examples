package slice_map_bm

import (
	crand "crypto/rand"
	"fmt"
	mrand "math/rand"
	"runtime"
	"testing"
	"time"
)

const (
	strLenMin = 5
	strLenMax = 50

	sliceSeekLimit = 20000
)

var (
	testSizes []int

	sliceItems []Data
	mapItems   map[string]Data
)

func BenchmarkSliceDyn(b *testing.B) {
	for _, size := range testSizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			start := time.Now()
			sliceItems = sliceFillDyn(size)
			reportExtraTime(b, start, "fill_sec/op")

			defer func() {
				start = time.Now()
				sliceItems = nil
				runtime.GC()
				reportExtraTime(b, start, "gc_sec/op")
			}()

			if size > sliceSeekLimit {
				return
			}

			start = time.Now()
			for s := 0; s < size; s++ {
				i := sliceSeek(sliceItems, allStrings[s])
				_ = i
			}
			reportExtraTime(b, start, "seek_sec/op")
		})
	}
}

func BenchmarkSliceFix(b *testing.B) {
	for _, size := range testSizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			start := time.Now()
			sliceItems = sliceFillFix(size)
			reportExtraTime(b, start, "fill_sec/op")

			defer func() {
				start = time.Now()
				sliceItems = nil
				runtime.GC()
				reportExtraTime(b, start, "gc_sec/op")
			}()

			if size > sliceSeekLimit {
				return
			}

			start = time.Now()
			for s := 0; s < size; s++ {
				i := sliceSeek(sliceItems, allStrings[s])
				_ = i
			}
			reportExtraTime(b, start, "seek_sec/op")

			start = time.Now()
			sliceItems = nil
			runtime.GC()
			reportExtraTime(b, start, "gc_sec/op")
		})
	}
}

func BenchmarkMapDyn(b *testing.B) {
	for _, size := range testSizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			start := time.Now()
			mapItems = mapFillDyn(size)
			reportExtraTime(b, start, "fill_sec/op")

			start = time.Now()
			for s := 0; s < size; s++ {
				data, _ := mapSeek(mapItems, allStrings[s])
				_ = data
			}
			reportExtraTime(b, start, "seek_sec/op")

			start = time.Now()
			mapItems = nil
			runtime.GC()
			reportExtraTime(b, start, "gc_sec/op")
		})
	}
}

func BenchmarkMapFix(b *testing.B) {
	for _, size := range testSizes {
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			start := time.Now()
			mapItems = mapFillFix(size)
			reportExtraTime(b, start, "fill_sec/op")

			start = time.Now()
			for s := 0; s < size; s++ {
				data, _ := mapSeek(mapItems, allStrings[s])
				_ = data
			}
			reportExtraTime(b, start, "seek_sec/op")

			start = time.Now()
			mapItems = nil
			runtime.GC()
			reportExtraTime(b, start, "gc_sec/op")
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

func reportExtraTime(b *testing.B, start time.Time, metric string) {
	b.ReportMetric(float64(time.Since(start).Nanoseconds())/float64(b.N), metric)
}
