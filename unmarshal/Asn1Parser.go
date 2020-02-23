package unmarshal

import (
	"bytes"
)

func Asn1ParserBuffer(buffer bytes.Buffer) *Asn1ParseResult {
	reader := NewAsn1Reader(buffer)
	return Asn1ParserReader(reader)
}

func Asn1ParserContainer(container *Asn1ParseResult) {
	reader := NewAsn1Reader(container.Buf)
	pos := container.BodyStart
	for {
		reader.Position = pos
		asn1Obj := Asn1ParserReader(reader)
		if asn1Obj == nil {
			break
		}

		container.AddElement(asn1Obj)
		tmplen := asn1Obj.GetEncodingLength()
		pos = pos + tmplen
		if asn1Obj.Tag.IsEOC() {
			break
		}

		if container.CheckBodyFinished(pos) {
			break
		}
	}
	container.BodyEnd = pos
}

func Asn1ParserReader(reader *Asn1Reader) *Asn1ParseResult {
	if reader.Buffer.Len() > reader.Buffer.Cap() {
		return nil
	}
	header := reader.ReadHeader()
	tmpTag := header.Tag
	bodyStart := reader.Position

	var parseResult *Asn1ParseResult
	conf := tmpTag.IsPrimitive()
	if conf {
		parseResult = NewAsn1ParseResult(*header, bodyStart, reader.Buffer)
	} else {
		parseResult = NewAsn1ParseResult(*header, bodyStart, reader.Buffer)
		length := header.Length
		if length != 0 {
			Asn1ParserContainer(parseResult)
		}
	}

	return parseResult
}
