package for_closure //nolint:golint,nosnakecase // underscore in package name

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoClosure_1_21(t *testing.T) {
	if !strings.HasPrefix(runtime.Version(), "go1.21.") {
		t.Skip("Not a Go 1.21")

		return
	}

	results := doClosure([]string{"a", "b", "c"})
	fmt.Println(results)

	stat := map[string]int{}
	for _, v := range results {
		stat[v]++
	}

	assert.Greater(t, stat["c"], 1)
}

func TestDoClosure_1_22(t *testing.T) {
	if !strings.HasPrefix(runtime.Version(), "go1.22.") {
		t.Skip("Not a Go 1.22")

		return
	}
	results := doClosure([]string{"a", "b", "c"})
	fmt.Println(results)

	stat := map[string]int{}
	for _, v := range results {
		stat[v]++
	}

	expected := map[string]int{"a": 1, "b": 1, "c": 1}
	assert.Equal(t, expected, stat)
}
