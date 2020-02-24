package assign

import (
	"github.com/xzy256/hasparse/unmarshal"
)

/**
LastReq ::= SEQUENCE OF SEQUENCE {
	lr-type         [0] Int32,
	lr-value        [1] KerberosTime
}
*/
type LastReq struct {
	Flag             *unmarshal.HasTag // 1
	ExplicitFieldMap map[int]interface{}

	Name             string
}

func (this *LastReq) SetName(name string) {
	this.Name = name
}

func (this *LastReq) Construct() {
	this.Flag = &unmarshal.HasTag{TagNo: 1}
	this.ExplicitFieldMap = make(map[int]interface{})
	this.ExplicitFieldMap[0] = &Asn1Integer{TagNo: 0x02}
	this.ExplicitFieldMap[1] = &Asn1GeneralizedTime{TagNo: 24}
}

func (this *LastReq) DecodeBody(result *unmarshal.Asn1ParseResult) {
	body := result.Children[0]
	for body.GetIndex() != len(this.ExplicitFieldMap) {
		body = body.Children[0]
	}
	for i := 0; i < body.GetIndex(); i++ {
		value := this.ExplicitFieldMap[i]
		tmpResult := body.Children[i]
		switch tmpResult.Tag.TagNo {
		case 0:
			intValue, ok := value.(*Asn1Integer)
			if ok {
				intValue.SetName("Last-req.lr-type[int]")
				intValue.DecodeBody(tmpResult.Children[0])
			}
		case 1:
			timeValue, ok := value.(*Asn1GeneralizedTime)
			if ok {
				timeValue.SetName("Last-req.lr-value[generalized time]")
				timeValue.DecodeBody(tmpResult.Children[0])
			}
		}
	}
}
