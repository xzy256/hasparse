package assign

import (
	"github.com/xzy256/hasparse/unmarshal"
)

const (
	NAME_TYPE   int = 0
	NAME_STRING int = 1
)

/*
PrincipalName   ::= SEQUENCE {
     name-type       [0] Int32,
     name-string     [1] SEQUENCE OF KerberosString
}
*/
type PrincipalName struct {
	NameType         int
	NameString       string

	Name 			 string
	Flag             *unmarshal.HasTag
	ExplicitFieldMap map[int]interface{}
}

func (this *PrincipalName) MakeSalt(realm string) string {
	return realm + this.NameString
}

func (this *PrincipalName) SetName(name string) {
	this.Name = name
}

func (this *PrincipalName) Construct() {
	this.Flag = &unmarshal.HasTag{}
	this.ExplicitFieldMap = make(map[int]interface{})
	this.ExplicitFieldMap[0] = &Asn1Integer{TagNo: 0x02}
	this.ExplicitFieldMap[1] = &Asn1OctString{}
}

func (this *PrincipalName) DecodeBody(result *unmarshal.Asn1ParseResult) {
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
				intValue.SetName("PrincipalName.name-type [int]")
				intValue.DecodeBody(tmpResult.Children[0])
			}
		case 1:
			octStrValue, ok := value.(*Asn1OctString)
			if ok {
				octStrValue.SetName("PrincipalName.name-string [oct string]")
				octStrValue.DecodeBody(tmpResult.Children[0])
			}
		}
	}
}
