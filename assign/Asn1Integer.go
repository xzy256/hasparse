package assign

import (
	"hasparse/unmarshal"
	"hasparse/utils"
)

type Asn1Integer struct {
	ValueBytes []byte // multi bytes for int32
}

func (this *Asn1Integer) Init() {
}

func (this *Asn1Integer) DecodeBody(parseResult *unmarshal.Asn1ParseResult) {
	remainingBytes := parseResult.GetBodyBuffer().Bytes()
	if len(remainingBytes) > 0 {
		this.ValueBytes = remainingBytes
	}
}

func (this *Asn1Integer) Value() int32 {
	if this == nil {
		return 0
	}
	return utils.BytesToInt(this.ValueBytes)
}
