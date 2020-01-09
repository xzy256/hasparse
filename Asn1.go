package main

import "bytes"

type Asn1 struct {
}

func NewAsn1() *Asn1 {
	return &Asn1{}
}

func (this *Asn1) EncodeBuffer(buffer bytes.Buffer, value Asn1Type) {
	value.DcodeBuffer(buffer)
}

/*func (this *Asn1) Encode(value Asn1Type) []byte {
	return value.Encode()
}*/

/*func (this *Asn1) DecodeBytes(context []byte) Asn1Type {
	buffer := bytes.NewBuffer(context)
	return 	this.DecodeBuffer(*buffer)
}*/

/*func (this *Asn1) DecodeBuffer(buffer bytes.Buffer) Asn1Type {
	parseResult := Asn1Parser.parse(buffer)
	return Asn1Converter.convert(parseResult, false)
}*/

//func (this *Asn1) ParseBuffer(content bytes.Buffer) Asn1ParseResult {
//	return
//}
