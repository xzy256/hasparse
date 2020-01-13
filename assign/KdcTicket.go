package assign

const (
	TKT_VNO     int = 0
	REALM       int = 1
	SNAME       int = 2
	TK_ENC_PART int = 3
)

type KdcTicket struct {
	Tktvno   int
	Realm    *Asn1OctetString
	Sname    *PrincipalName
	EncPart  *EncryptedData
	filedInfo []int
	position []byte
}

func (this *KdcTicket) Init() {
	this.position = []byte{5, 0, 0}
	this.filedInfo = []int{TKT_VNO, REALM, SNAME, TK_ENC_PART}
}
