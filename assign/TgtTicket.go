package assign

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"gopkg.in/jcmturner/gokrb5.v7/crypto"
	krbcom "gopkg.in/jcmturner/gokrb5.v7/crypto/common"
	"gopkg.in/jcmturner/gokrb5.v7/crypto/rfc3962"
	"gopkg.in/jcmturner/gokrb5.v7/iana/keyusage"
	"gopkg.in/jcmturner/gokrb5.v7/types"
	"hasparse/utils"
)

func HandleKdcRep(kdcRep *KdcRep, passPhrass string) {
	//etype := EncryptionTypeFromValue(kdcRep.EncData.Etype)
	//salt := kdcRep.Cname.MakeSalt(kdcRep.Crealm)
	//clientKey := String2Key(passPhrass, salt, etype)

	decryptWith(kdcRep)

	//plainText := Decrypt(kdcRep.EncData.Cipher.ValueBytes, clientKey.Keyvalue.ValueBytes)
	//fmt.Println("++++", plainText)
}

func decryptWith(kdcRep *KdcRep) {
	constant := []byte{0,0,0,3,0} // 3-->AS_REP_ENCPART
	constant[4] = byte(0xaa)

	ki := Selfnfold(constant, aes.BlockSize)
	fmt.Println(ki)

	u := krbcom.GetUsageKe(3)
	fmt.Println(u)

	key := &types.EncryptionKey{
		KeyType:17,
		KeyValue:ki,
	}
	plain, _ := Decrypt2(kdcRep, *key)
	fmt.Println(plain)


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

func Decrypt(data []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func Decrypt2(kdcRep *KdcRep, key types.EncryptionKey) ([]byte, error) {
	ticketD := &types.EncryptedData{
		EType:int32(kdcRep.EncData.Etype),
		KVNO: int(kdcRep.EncData.Kvno),
		Cipher:kdcRep.EncData.Cipher.ValueBytes,
	}
	b, err := crypto.DecryptEncPart(*ticketD, key, keyusage.AS_REP_ENCPART)
	if err != nil {
		return nil, fmt.Errorf("error decrypting Ticket EncPart: %v", err)
	}

	return b, nil
}
