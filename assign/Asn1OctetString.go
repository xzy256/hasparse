package assign

import (
	"hasparse/unmarshal"
)

type Asn1OctetString struct {
	ValueBytes []byte
}

func (this *Asn1OctetString) Init(){
}

func (this *Asn1OctetString)DecodeBody(parseResult *unmarshal.Asn1ParseResult) {
	remainingBytes := parseResult.GetBodyBuffer().Bytes()
	if len(remainingBytes) > 0 {
		this.ValueBytes = remainingBytes
	}
}

func (this *Asn1OctetString)Value() string {
	if this != nil {
		return string(this.ValueBytes)

	}
	return ""
}

