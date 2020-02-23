package assign

import (
	"hasparse/unmarshal"
)

type Asn1String struct {
	TagNo      int

	ValueBytes []byte
	Name       string
}

func (this *Asn1String) SetName(name string) {
	this.Name = name
}

func (this *Asn1String) Init() {
}

func (this *Asn1String) DecodeBody(parseResult *unmarshal.Asn1ParseResult) {
	body := parseResult
	for body.GetIndex() != 0 {
		body = body.Children[0]
	}
	remainingBytes := body.GetBodyBuffer().Bytes()
	if len(remainingBytes) > 0 {
		this.ValueBytes = remainingBytes
	}
}

func (this *Asn1String) Value() string {
	if this != nil {
		return string(this.ValueBytes)

	}
	return ""
}
