package slice_map_bm

const (
	nMin = 2  // 4
	nMax = 20 // 1048576
	iMax = 1 << nMax
)

type Data struct {
	str string
}

var (
	allStrings = [iMax]string{}
)

func sliceFillDyn(n int) []Data {
	items := make([]Data, 0)
	for i := 0; i < n; i++ {
		items = append(items, Data{allStrings[i]})
	}

	return items
}

func sliceFillFix(n int) []Data {
	items := make([]Data, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, Data{allStrings[i]})
	}

	return items
}

func sliceSeek(items []Data, seek string) int {
	for i, item := range items {
		if seek == item.str {
			return i
		}
	}

	return -1
}

func mapFillDyn(n int) map[string]Data {
	items := make(map[string]Data)
	for i := 0; i < n; i++ {
		items[allStrings[i]] = Data{allStrings[i]}
	}

	return items
}

func mapFillFix(n int) map[string]Data {
	items := make(map[string]Data, n)
	for i := 0; i < n; i++ {
		items[allStrings[i]] = Data{allStrings[i]}
	}

	return items
}

func mapSeek(items map[string]Data, seek string) (Data, bool) {
	data, has := items[seek]
	return data, has
}
