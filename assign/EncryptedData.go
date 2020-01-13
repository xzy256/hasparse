package assign

const (
	ETYPE  int = 0
	KVNO   int = 1
	CIPHER int = 2
)

type EncryptedData struct {
	Etype     int
	Kvno      uint32 //optional
	Cipher    *Asn1OctetString
	filedInfo []int
	position  []byte
}

func (this *EncryptedData) Init() {
	this.position = []byte{6, 0, 0}
	this.filedInfo = []int{ETYPE, KVNO, CIPHER}
}
