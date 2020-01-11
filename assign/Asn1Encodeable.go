package assign

import (
	"bytes"
	"hasparse/common"
	"hasparse/unmarshal"
	"log"
)

type Asn1Encodeable struct {
	BodyLength      int // -1
	OuterEncodeable *Asn1Encodeable
	EncodingT       common.EncodingType // BER
	IsImplct        bool                //true
	IsDLength       bool                //true
	Tag             *unmarshal.HasTag
}

func NewAsn1Encodeable(tag *unmarshal.HasTag) *Asn1Encodeable {
	return &Asn1Encodeable{
		BodyLength: -1,
		EncodingT:  common.BER,
		IsImplct:   true,
		IsDLength:  true,
		Tag:        tag,
	}
}

func NewAsn1EncodeableInt(tag common.UniversalTag) *Asn1Encodeable {
	return &Asn1Encodeable{
		BodyLength: -1,
		EncodingT:  common.BER,
		IsImplct:   true,
		IsDLength:  true,
		Tag:        unmarshal.NewHasUniversalTag(tag),
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
	this.EncodingT = common.BER
}

func (this *Asn1Encodeable) IsBER() bool {
	return this.EncodingT == common.BER
}

func (this *Asn1Encodeable) DcodeBytes(context []byte) {
	buf := bytes.NewBuffer(context)
	this.DcodeBuffer(*buf)
}

func (this *Asn1Encodeable) DcodeBuffer(buffer bytes.Buffer) {
	parseResult := unmarshal.Asn1ParserBuffer(buffer)
	this.DcodeAsn1ParseResult(parseResult)
}

func (this *Asn1Encodeable) DcodeAsn1ParseResult(parseResult *unmarshal.Asn1ParseResult) {
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

func (this *Asn1Encodeable) TaggedDecode(parseResult unmarshal.Asn1ParseResult, taggingOption TaggingOption) {
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
