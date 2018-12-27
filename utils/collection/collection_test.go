package collection

import (
	"github.com/golang-collections/go-datastructures/bitarray"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitArrayToByteArray(t *testing.T) {
	bits := bitarray.NewBitArray(64)
	bits.SetBit(7)
	bits.SetBit(14)
	bits.SetBit(22)
	bits.SetBit(23)
	expectedBytes := []byte{1, 2, 3, 0, 0, 0, 0, 0}
	bytes := BitArrayToByteArray(bits)
	assert.Equal(t, expectedBytes, bytes)
}
