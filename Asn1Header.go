package main

type Asn1Header struct {
	Tag    HasTag
	Length int
}

func NewAsn1Header(tag HasTag, length int) *Asn1Header {
	return &Asn1Header{
		Tag:    HasTag{tag.TagNo, tag.TagFlags},
		Length: length,
	}
}
