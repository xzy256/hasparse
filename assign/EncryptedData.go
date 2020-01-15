package assign

const (
	ETYPE  int = 0
	KVNO   int = 1
	CIPHER int = 2
)

/*
EncryptedData   ::= SEQUENCE {
	etype   [0] Int32 -- EncryptionType --,
	kvno    [1] UInt32 OPTIONAL,
	cipher  [2] OCTET STRING -- ciphertext
}
*/

type EncryptedData struct {
	Etype    int
	Kvno     uint32 //optional
	Cipher   *Asn1OctetString
	Optional bool
}
