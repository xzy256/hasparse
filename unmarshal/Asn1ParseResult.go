package unmarshal

import (
	"bytes"
	"hasparse/utils"
)

type Asn1ParseResult struct {
	Tag       HasTag
	Header    Asn1Header
	BodyStart int
	BodyEnd   int
	Buf       bytes.Buffer
	Children  []*Asn1ParseResult
	index     int
}

func (this *Asn1ParseResult) GetBodyBuffer() *bytes.Buffer {
	if this.BodyEnd >= this.BodyStart {
		return bytes.NewBuffer(this.Buf.Bytes()[this.BodyStart:this.BodyEnd])
	}
	return bytes.NewBuffer(this.Buf.Bytes()[this.BodyStart:])
}

func (this *Asn1ParseResult) GetElementByIndex(index int) *Asn1ParseResult {
	return this.Children[index]
}

func (this *Asn1ParseResult) AddElement(newAsn1Obj *Asn1ParseResult) {
	this.Children[this.index] = newAsn1Obj
	this.index ++
}

func NewAsn1ParseResult(header Asn1Header, bodyStart int, buf bytes.Buffer) *Asn1ParseResult {
	bodyEnd := bodyStart + header.Length
	if header.Length == -1 {
		bodyEnd = -1
	}

	return &Asn1ParseResult{
		Tag:       header.Tag,
		Header:    header,
		BodyStart: bodyStart,
		BodyEnd:   bodyEnd,
		Buf:       buf,
		Children:  make([]*Asn1ParseResult, buf.Len()),
		index:     0,
	}
}

func (this *Asn1ParseResult) GetBodyLength() int {
	if this.Header.Length != -1 {
		return this.Header.Length
	} else if this.BodyEnd != -1 {
		return this.BodyEnd - this.BodyStart
	}
	return -1
}

func (this *Asn1ParseResult) GetHeaderLength() int {
	bodyLen := -1
	if this.Header.Length != -1 {
		bodyLen = this.Header.Length
	} else if this.BodyEnd != -1 {
		bodyLen = this.BodyEnd - this.BodyStart
	}

	headerLen := utils.LengthOfTagLength(this.Header.Tag.TagNo)
	if this.Header.Length != -1 {
		headerLen += utils.LengthOfBodyLength(bodyLen)
	} else {
		headerLen += 1
	}

	return headerLen
}

func (this *Asn1ParseResult) GetOffset() int {
	return this.BodyStart - this.Header.Length
}

func (this *Asn1ParseResult) CheckBodyFinished(pos int) bool {
	return this.BodyEnd != -1 && pos >= this.BodyEnd
}

func (this *Asn1ParseResult) GetEncodingLength() int {
	headerLen := this.GetHeaderLength()
	bodyLen := this.GetBodyLength()
	return headerLen + bodyLen
}
