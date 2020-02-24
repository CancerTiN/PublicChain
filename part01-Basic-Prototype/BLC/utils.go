package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Convert int64 to []byte
func Int64ToHex(data int64) (byteSlice []byte) {
	w := new(bytes.Buffer)
	err := binary.Write(w, binary.BigEndian, data)
	if err != nil {
		log.Panic(err)
	}
	byteSlice = w.Bytes()
	return
}

