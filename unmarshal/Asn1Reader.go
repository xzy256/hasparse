package unmarshal

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

type Asn1Reader struct {
	Buffer   bytes.Buffer
	Position int
}

func NewAsn1Reader(buffer bytes.Buffer) *Asn1Reader {
	return &Asn1Reader{
		Buffer: buffer,
		Position: 0,
	}
}

func (this *Asn1Reader) GetLength() int {
	return this.Buffer.Len()
}

func (this *Asn1Reader) ReadByte() byte {
	b := this.Buffer.Bytes()[this.Position]
	this.Position++
	return b
}

func (this *Asn1Reader) GetByte() byte {
	return this.Buffer.Bytes()[this.Position]
}

func (this *Asn1Reader) ReadHeader() *Asn1Header {
	tag := this.ReadTag()
	if this.Position == 66 {
		fmt.Println("=====", tag)
	}
	valueLength := this.ReadLength()
	header := NewAsn1Header(*tag, valueLength)

	return header
}

func (this *Asn1Reader) ReadTag() *HasTag {
	tagFlags := this.ReadTagFlags()
	tagNo := this.ReadTagNo(tagFlags)
	return NewHasTagFlag(tagFlags, tagNo)
}

func (this *Asn1Reader) ReadTagFlags() int {
	b := this.ReadByte()
	tagFlags := b & 0xff
	return int(tagFlags)
}

func (this *Asn1Reader) ReadTagNo(tagFlags int) int {
	tagNo := tagFlags & 0x1f
	if tagNo == 0x1f {
		tagNo = 0
		b := this.ReadByte()
		this.Position += 1
		b = b & 0xff
		if (b & 0x7f) == 0 {
			err := errors.New("invalid high tag number found")
			fmt.Println(err)
			os.Exit(-1)
		}

		for b >= 0 && (b&0x80) != 0 {
			tagNo |= int(b & 0x7f)
			tagNo <<= 7
			b = this.ReadByte()
			this.Position += 1
		}
		tagNo |= int(b & 0x7f)
	}

	return tagNo
}

func (this *Asn1Reader) ReadLength() int {
	t := this.ReadByte()

	result := t & 0xff
	if result == 0x80 {
		return -1
	}
	var res = int(result)
	if result > 127 {
		length := result & 0x7f
		if length > 4 {
			fmt.Println("bad length of more than 4 bytes: ", length)
			os.Exit(-1)
		}
		res = 0
		var tmp byte
		for i := 0; i < int(length); i++ {
			tmp = this.ReadByte()
			tmp = tmp & 0xff
			res = (res << 8) + int(tmp)
		}
	}

	if res < 0 {
		fmt.Println("Invalid length ", result)
		os.Exit(-1)
	}

	return res
}
