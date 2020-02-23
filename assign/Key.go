package assign

import (
	"hasparse/unmarshal"
)

type Key struct {
	TagNo            int // 0
	ExplicitFieldMap map[int]interface{}

	Name 			 string
}

func (this *Key) SetName(name string) {
	this.Name = name
}

func (this *Key) Construct() {
	this.TagNo = 0
	this.ExplicitFieldMap = make(map[int]interface{})
	this.ExplicitFieldMap[0] = &Asn1Integer{TagNo: 0x02}
	this.ExplicitFieldMap[1] = &Asn1OctString{TagNo: 0x04}
}

func (this *Key) DecodeBody(result *unmarshal.Asn1ParseResult) {
	body := result
	for body.GetIndex() != len(this.ExplicitFieldMap)  {
		body = body.Children[0]
	}

	for i := 0; i < body.GetIndex(); i++ {
		value := this.ExplicitFieldMap[i]
		tmpResult := body.Children[i]
		switch tmpResult.Tag.TagNo {
		case 0:
			intValue, ok := value.(*Asn1Integer)
			if ok {
				intValue.SetName("Key.type [int]")
				intValue.DecodeBody(tmpResult.Children[0])
			}
		case 1:
			octStrValue, ok := value.(*Asn1OctString)
			if ok {
				octStrValue.SetName("key.value [oct string]")
				octStrValue.DecodeBody(tmpResult.Children[0])
			}
		}
	}
}
