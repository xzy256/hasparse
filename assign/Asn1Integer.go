package assign

import (
	"github.com/xzy256/hasparse/unmarshal"
	"github.com/xzy256/hasparse/utils"
)

type Asn1Integer struct {
	TagNo      int
	ValueBytes []byte // multi bytes for int32

	Name       string
}

func (this *Asn1Integer) Init() {
	this.TagNo = 2
}

func (this *Asn1Integer) SetName(name string) {
	this.Name = name
}

func (this *Asn1Integer) DecodeBody(parseResult *unmarshal.Asn1ParseResult) {
	body := parseResult
	for body.GetIndex() != 0 {
		body = body.Children[0]
	}
	remainingBytes := body.GetBodyBuffer().Bytes()
	if len(remainingBytes) > 0 {
		this.ValueBytes = remainingBytes
	}
}

func (this *Asn1Integer) Value() int32 {
	if this == nil {
		return 0
	}
	return utils.BytesToInt(this.ValueBytes, false)
}
