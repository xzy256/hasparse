package assign

/**
EncryptionKey   ::= SEQUENCE {
	keytype         [0] Int32 -- actually encryption type --,
	keyvalue        [1] OCTET STRING
}
*/

type EncryptionKey struct {
	Keytype  *EncryptionType
	Keyvalue *Asn1OctetString
	Kvno     int
}

func NewEncryptionKey(keytype *EncryptionType, keyvalue []byte) *EncryptionKey {
	return &EncryptionKey{
		Keytype:  keytype,
		Keyvalue: &Asn1OctetString{ValueBytes: keyvalue},
		Kvno:     -1,
	}
}
