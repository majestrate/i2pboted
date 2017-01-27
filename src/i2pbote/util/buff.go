package util

import "encoding/binary"

// get as uint16 but int
func UInt16_i(data []byte) int {
	return int(binary.BigEndian.Uint16(data))
}

// get as uint16
func UInt16(data []byte) uint16 {
	return binary.BigEndian.Uint16(data)
}

// get as uint32
func UInt32(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

func PutUInt16_i(i int, data []byte) {
	binary.BigEndian.PutUint16(data[:], uint16(i))
}
