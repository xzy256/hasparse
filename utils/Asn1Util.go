package utils

import (
	"encoding/binary"
	"log"
)

func LengthOfTagLength(tagNo int) int {
	length := 1
	if tagNo >= 31 {
		if tagNo < 128 {
			length++
		} else {
			length++
			var condition = true
			for condition || tagNo > 127 {
				condition = false
				tagNo >>= 7
				length++
			}
		}
	}

	return length
}

func LengthOfBodyLength(bodyLength int) int {
	length := 1
	if bodyLength > 127 {
		payload := bodyLength
		for payload != 0 {
			payload >>= 8
			length++
		}
	}

	return length
}

func IntToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func BytesToInt(buf []byte) int32 {
	if len(buf) > 4 {
		log.Fatal("Array out of bounds, 4 bytes for int32, your len:", len(buf))
	}
	length := len(buf)
	tmpbytes := []byte{0, 0, 0, 0}
	for i := 0; i < length; i++ {
		tmpbytes[3-i] = buf[i]
	}
	return int32(binary.BigEndian.Uint32(tmpbytes))
}
