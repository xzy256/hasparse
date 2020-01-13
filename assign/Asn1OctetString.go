package assign

import (
	"hasparse/unmarshal"
)

type Asn1OctetString struct {
	ValueBytes []byte
	filedInfo  []int
	position   []byte
}

func (this *Asn1OctetString) Init(){
	this.position = []byte{1, 0, 0}
	this.filedInfo = []int{1}
}

func (this *Asn1OctetString)DecodeBody(parseResult *unmarshal.Asn1ParseResult) {
	remainingBytes := parseResult.GetBodyBuffer().Bytes()
	if len(remainingBytes) > 0 {
		this.ValueBytes = remainingBytes
	}
}

func (this *Asn1OctetString)Value() string {
	return string(this.ValueBytes)
}

