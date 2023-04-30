package enum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnum(t *testing.T) {
	//x := NewEnum[int](2)
	x := ""

	a := statusInvalid
	b := statusAvailable

	assert.Equal(t, "x", fmt.Sprintf("%v %v %s", a, b, x))
}
