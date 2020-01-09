package main

import (
	"bytes"
	"log"
)

type Asn1Encodeable struct {
	BodyLength      int // -1
	OuterEncodeable *Asn1Encodeable
	EncodingT       EncodingType // BER
	IsImplct        bool         //true
	IsDLength       bool         //true
	Tag             *HasTag
}

func NewAsn1Encodeable(tag *HasTag) *Asn1Encodeable {
	return &Asn1Encodeable{
		BodyLength: -1,
		EncodingT:  BER,
		IsImplct:   true,
		IsDLength:  true,
		Tag:        tag,
	}
}

func NewAsn1EncodeableInt(tag UniversalTag) *Asn1Encodeable {
	return &Asn1Encodeable{
		BodyLength: -1,
		EncodingT:  BER,
		IsImplct:   true,
		IsDLength:  true,
		Tag:        NewHasUniversalTag(tag),
	}
}

func (this *Asn1Encodeable) IsPrimitive() bool {
	return this.Tag.IsPrimitive()
}

func (this *Asn1Encodeable) UsePrimitive(isP bool) {
	this.Tag.UsePrimitive(isP)
}

func (this *Asn1Encodeable) UseDefinitiveLength(isdl bool) {
	this.IsDLength = isdl
}

func (this *Asn1Encodeable) IsDefinitiveLength() bool {
	return this.IsDLength
}

func (this *Asn1Encodeable) UseImplicit(isImplicit bool) {
	this.IsImplct = isImplicit
}

func (this *Asn1Encodeable) IsImplicit() bool {
	return this.IsImplct
}

func (this *Asn1Encodeable) UseBER() {
	this.EncodingT = BER
}

func (this *Asn1Encodeable) IsBER() bool {
	return this.EncodingT == BER
}

func (this *Asn1Encodeable) DcodeBytes(context []byte) {
	buf := bytes.NewBuffer(context)
	this.DcodeBuffer(*buf)
}

func (this *Asn1Encodeable) DcodeBuffer(buffer bytes.Buffer) {
	parseResult := Asn1ParserBuffer(buffer)
	this.DcodeAsn1ParseResult(parseResult)
}

func (this *Asn1Encodeable) DcodeAsn1ParseResult(parseResult *Asn1ParseResult) {
	//tmpParseResult := parseResult
	tag1 := this.Tag
	tag2 := parseResult.Tag
	if !tag1.Equal(&tag2) { // Primitive but using constructed encoding
		if this.IsPrimitive() && !parseResult.Tag.IsPrimitive() {
			// TODO
		} else {
			log.Fatal("unexpected item:", parseResult.Tag, "expecting:", this.Tag)
		}
	}
	//decodeBody(tmpParseResult)
}

func (this *Asn1Encodeable) TaggedDecode(parseResult Asn1ParseResult, taggingOption TaggingOption) {
	expectTaggingTagFlags := taggingOption.GetTag(!this.IsPrimitive())
	//tmpParseResult := parseResult
	if !expectTaggingTagFlags.Equal(&parseResult.Tag) {
		if this.IsPrimitive() && !parseResult.Tag.IsPrimitive() {
			// TODO
		} else {
			log.Fatal("unexpected item:", parseResult.Tag, "expecting:", expectTaggingTagFlags)
		}
	}

	if taggingOption.IsImplicit {
		//decodeBody(tmpParseResult)
	}else{
		//tmpParseResult := parseResult.Children.Front()
	}
}
