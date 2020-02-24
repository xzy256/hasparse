package assign

import (
	"github.com/xzy256/hasparse/unmarshal"
)

/**
EncKDCRepPart   ::= SEQUENCE {
	key             [0] Key,
	last-req        [1] LastReq,
	nonce           [2] UInt32,
	key-expiration  [3] KerberosTime OPTIONAL,
	flags           [4] TicketFlags,
	authtime        [5] KerberosTime,
	starttime       [6] KerberosTime OPTIONAL,
	endtime         [7] KerberosTime,
	renew-till      [8] KerberosTime OPTIONAL,
	srealm          [9] Realm,
	sname           [10] PrincipalName,
	caddr           [11] HostAddresses OPTIONAL
}
*/

type EncKdcRepPart struct {
	Flag             *unmarshal.HasTag // 25
	ExplicitFieldMap map[int]interface{}
}

func (this *EncKdcRepPart) Construct() {
	this.Flag = &unmarshal.HasTag{TagNo: 25}
	this.ExplicitFieldMap = make(map[int]interface{})
	key := &Key{}
	key.Construct()
	this.ExplicitFieldMap[0] = key // key
	lastReq := &LastReq{}
	lastReq.Construct()
	this.ExplicitFieldMap[1] = lastReq                           //last-req
	this.ExplicitFieldMap[2] = &Asn1Integer{TagNo: 0x02}         // nonce
	this.ExplicitFieldMap[3] = &Asn1GeneralizedTime{TagNo: 0x18} // key-expiration
	this.ExplicitFieldMap[4] = &Asn1Flags{TagNo: 0x3}            // flags
	this.ExplicitFieldMap[5] = &Asn1GeneralizedTime{TagNo: 0x18} // at
	this.ExplicitFieldMap[6] = &Asn1GeneralizedTime{TagNo: 0x18} // st
	this.ExplicitFieldMap[7] = &Asn1GeneralizedTime{TagNo: 0x18} // et
	this.ExplicitFieldMap[8] = &Asn1GeneralizedTime{TagNo: 0x18} // rt
	this.ExplicitFieldMap[9] = &Asn1String{TagNo: 0x1B}
	princ := &PrincipalName{}
	princ.Construct()
	this.ExplicitFieldMap[10] = princ
}

func (this *EncKdcRepPart) DecodeBody(parseResult *unmarshal.Asn1ParseResult) {
	var index int
	for i := 0; i < parseResult.GetIndex(); i++ {
		index = parseResult.Children[i].Tag.TagNo
		value := this.ExplicitFieldMap[index]
		switch index {
		case 0:
			keyVlue, ok := value.(*Key)
			if ok {
				keyVlue.SetName("Key [sequence]")
				keyVlue.DecodeBody(parseResult.Children[i])
			}
		case 1:
			lastReq, ok := value.(*LastReq)
			if ok {
				lastReq.SetName("Last-req [sequence]")
				lastReq.DecodeBody(parseResult.Children[i])
			}
		case 2:
			intValue, ok := value.(*Asn1Integer)
			if ok {
				intValue.SetName("nonce [int]")
				intValue.DecodeBody(parseResult.Children[i])
			}
		case 4:
			krbFlag, ok := value.(*Asn1Flags)
			if ok {
				krbFlag.SetName("flags [bit string]")
				krbFlag.DecodeBody(parseResult.Children[i])
			}
		case 3:
			fallthrough
		case 5:
			fallthrough
		case 6:
			fallthrough
		case 7:
			fallthrough
		case 8:
			gtime, ok := value.(*Asn1GeneralizedTime)
			if ok {
				gtime.SetName("generalized time")
				gtime.DecodeBody(parseResult.Children[i])
			}
		case 9:
			srealm, ok := value.(*Asn1String)
			if ok {
				srealm.SetName("srealm [string]")
				srealm.DecodeBody(parseResult.Children[i])
			}
		case 10:
			princp, ok := value.(*PrincipalName)
			if ok {
				princp.SetName("sname [sequence]")
				princp.DecodeBody(parseResult.Children[i])
			}
		}
	}
}
