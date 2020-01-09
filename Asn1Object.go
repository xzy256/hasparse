package main

type Asn1Object struct {
	Tag 	*HasTag
}

func NewAsn1Object(tag *HasTag) *Asn1Object{
	return &Asn1Object{
		Tag:NewHasTagFlag(tag.TagFlags, tag.TagNo),
	}
}

