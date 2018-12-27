package collection

import (
	"github.com/golang-collections/go-datastructures/bitarray"
	"math"
)

/*
BitArrayToByteArray converts an array of bits to an array of bytes, e.g.,
|0000|0001|b to |0|8|h.
*/
func BitArrayToByteArray(bits bitarray.BitArray) []byte {
	bytes := make([]byte, sizeInBitsToBytes(bits.Capacity()))
	for i := 0; i < int(bits.Capacity()); i++ {
		bit, _ := bits.GetBit(uint64(i))
		if bit {
			bytes[i>>3] = bytes[i>>3] | byte(1<<(7-(uint(i)&0x07)))
		}
	}
	return bytes
}

func sizeInBitsToBytes(sizeInBits uint64) int {
	return int(math.Ceil(float64(sizeInBits / 8)))
}
