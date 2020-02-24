package assign

import (
	"github.com/xzy256/hasparse/unmarshal"
)

func DecodeBody(parseResult *unmarshal.Asn1ParseResult, class interface{}) {
	encKdcRp, ok := class.(EncKdcRepPart)
	if ok {
		encKdcRp.DecodeBody(parseResult.Children[0])
	}

}
