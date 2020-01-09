package main

import (
	"bytes"
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

func EncodeTag(buffer bytes.Buffer, tag *HasTag) {
	flags := tag.TagFlags
	tagNo := tag.TagNo

	if tagNo < 31 {
		buffer.WriteByte(byte(flags | tagNo))
	} else {
		buffer.WriteByte(byte(flags | 0x1f))
		if tagNo < 128 {
			buffer.WriteByte(byte(tagNo))
		} else {
			tmpBytes := make([]byte, 5)
			iPut := len(tmpBytes)
			tmpBytes[iPut-1] = (byte)(tagNo & 0x7f)
			precond := true
			for precond || tagNo > 127 {
				precond = false
				tagNo >>= 7
				iPut--
				tmpBytes[iPut] = (byte)(tagNo&0x7f | 0x80)
			}
			buffer.Write(tmpBytes[iPut : buffer.Len()-iPut])
		}
	}
}

func EncodeLength(buffer bytes.Buffer, bodyLength int) {
	if bodyLength < 128 {
		buffer.WriteByte(byte(bodyLength))
	} else {
		length := 0
		payload := bodyLength
		for payload != 0 {
			payload >>= 8
			length++
		}
		buffer.WriteByte(byte(length | 0x80))
		payload = bodyLength
		for i := length - 1; i >= 0; i-- {
			buffer.WriteByte(byte(payload >> uint(i) * 8))
		}
	}
}
