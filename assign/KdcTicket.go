package assign

const (
	TKT_VNO     int = 0
	REALM       int = 1
	SNAME       int = 2
	TK_ENC_PART int = 3
)

/**
Ticket          ::= [APPLICATION 1] SEQUENCE {
	tkt-vno         [0] INTEGER (5),
	realm           [1] Realm,
	sname           [2] PrincipalName,
	enc-part        [3] EncryptedData -- EncTicketPart
}
*/
type KdcTicket struct {
	Tktvno   int
	Realm    string
	Sname    *PrincipalName
	EncPart  *EncryptedData
}

func (this *KdcTicket) Init() {
	this.Sname = &PrincipalName{}
	this.EncPart = &EncryptedData{}
}
