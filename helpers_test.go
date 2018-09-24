package smtcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToUint(t *testing.T) {
	raw := []byte{1, 0, 0, 0}
	expected := uint32(1)
	actual := BytesToUint(raw)

	assert.Equal(t, expected, actual)
}

func TestUint32ToBytes(t *testing.T) {
	raw := uint32(1)
	expected := []byte{1, 0, 0, 0}
	actual := Uint32ToBytes(raw)

	assert.Equal(t, expected, actual)
}

func TestUint32ToBytesAndBytesToUint32Integration(t *testing.T) {
	twoTo32 := 4294967296
	step := 5000
	for i := 0; i < twoTo32; i += step {
		raw := uint32(i)
		assert.Equal(t, raw, BytesToUint(Uint32ToBytes(raw)))
	}

	for bit1 := 0; bit1 < 255; bit1 += 32 {
		for bit2 := 0; bit2 < 255; bit2 += 32 {
			for bit3 := 0; bit3 < 255; bit3 += 32 {
				for bit4 := 0; bit4 < 255; bit4 += 32 {
					raw_bytes := []byte{0, 0, 0, 0}
					raw_bytes[0] = byte(bit1)
					raw_bytes[1] = byte(bit2)
					raw_bytes[2] = byte(bit3)
					raw_bytes[3] = byte(bit4)
					assert.Equal(t, raw_bytes, Uint32ToBytes(BytesToUint(raw_bytes)))
				}
			}
		}
	}
}
