package randx

import (
	"crypto/rand"
	"encoding/binary"
)

// Uint64 is used to generate a random uint64 number.
func Uint64() (uint64, error) {
	var ret uint64
	err := binary.Read(rand.Reader, binary.LittleEndian, &ret)
	if err != nil {
		return 0, err
	}

	return ret, nil
}
