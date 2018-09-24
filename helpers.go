package smtcp

import "encoding/binary"

func BytesToUint(b []byte) uint32 {
	if len(b) != 4 {
		return 0
	}
	return binary.LittleEndian.Uint32(b)
}

func Uint32ToBytes(v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)

	return b
}
