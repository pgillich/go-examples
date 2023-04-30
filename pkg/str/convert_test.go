package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString2ByteSlice(t *testing.T) {
	// given
	sampleStr1 := "012"
	// when
	sampleBytes1 := String2ByteSlice(sampleStr1)
	// then
	assert.Equal(t, []byte{0x30, 0x31, 0x32}, sampleBytes1)

	// when panic!
	sampleBytes1[1] = 0x39
}

func TestByteSlice2String(t *testing.T) {
	// given
	sampleBytes1 := []byte{0x30, 0x31, 0x32}
	// when
	sampleStr1 := ByteSlice2String(sampleBytes1)
	// then
	assert.Equal(t, "012", sampleStr1)

	// when
	sampleBytes1[1] = 0x39
	// then
	assert.Equal(t, "012", sampleStr1)
}
