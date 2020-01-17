package assign

import (
	"gopkg.in/jcmturner/gokrb5.v7/crypto"
	"gopkg.in/jcmturner/gokrb5.v7/crypto/rfc3962"
	"hasparse/utils"
)

func HandleKdcRep(kdcRep *KdcRep, passPhrass string) {
	//etype := EncryptionTypeFromValue(kdcRep.EncData.Etype)
	//salt := kdcRep.Cname.MakeSalt(kdcRep.Crealm)
	//clientKey := String2Key(passPhrass, salt, etype)

}

func String2Key(passPhrase string, salt string, etype *EncryptionType) *EncryptionKey {
	var e crypto.Aes128CtsHmacSha96
	s2kparams := utils.IterationsToS2Kparams(uint32(4096))
	keyBytes, err := rfc3962.StringToKey(passPhrase, salt, s2kparams, e)
	if err != nil {
		return nil
	}
	return NewEncryptionKey(etype, keyBytes)
}
