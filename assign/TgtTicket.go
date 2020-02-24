package assign

import (
	"bytes"
	"github.com/xzy256/hasparse/unmarshal"
	"github.com/xzy256/hasparse/utils"
	"gopkg.in/jcmturner/gokrb5.v7/crypto"
	"gopkg.in/jcmturner/gokrb5.v7/crypto/rfc3962"
)

func HandleKdcRep(kdcRep *KdcRep, passPhrass string) {
	etype := EncryptionTypeFromValue(kdcRep.EncData.Etype)
	salt := kdcRep.Cname.MakeSalt(kdcRep.Crealm)
	clientKey := String2Key(passPhrass, salt, etype)

	decryptedData := decryptWith(kdcRep, clientKey.Keyvalue.ValueBytes)
	if (decryptedData[0] & 0x1f) == 26 {
		decryptedData[0] = decryptedData[0] - 1
	}

	buf := bytes.NewBuffer(decryptedData)
	s1 := unmarshal.Asn1ParserBuffer(*buf)
	encKdcRp := EncKdcRepPart{}
	encKdcRp.Construct()
	DecodeBody(s1, encKdcRp)
	kdcRep.EncPart = &encKdcRp
}

func decryptWith(kdcRep *KdcRep, clientKey []byte) []byte{
	constant := []byte{0, 0, 0, 3, 0} // 3-->AS_REP_ENCPART
	constant[4] = byte(0xaa)

	//siv := Selfnfold(constant, aes.BlockSize)

	encProvider := &Aes128Provider{}
	encProvider.Construct()

	ke := Dr(clientKey, constant, encProvider)
	//constant[4] = byte(0x55)
	//ki := Dr(clientKey, constant, encProvider) // for check

	totalLen := len(kdcRep.EncData.Cipher.ValueBytes)
	aes128 := Aes128Provider{}
	confoundedLen := aes128.BlockSize()
	checksumLen := aes128.ChecksumSize()
	dataLen := totalLen - (confoundedLen + checksumLen)

	//workLens := []int{confoundedLen, checksumLen, dataLen}
	//fmt.Println(workLens)

	// decrypt and verify checksum
	//iv := make([]byte, encProvider.BlockSize())
	tmpEnc := make([]byte, confoundedLen+ dataLen)
	utils.ArrayCopy(kdcRep.EncData.Cipher.ValueBytes, 0, tmpEnc, 0, confoundedLen+ dataLen)
	raw := false
	if !raw {
		checksum := make([]byte, checksumLen)
		utils.ArrayCopy(kdcRep.EncData.Cipher.ValueBytes, confoundedLen+ dataLen, checksum, 0, checksumLen)
		output := make([]byte, len(tmpEnc))
		encProvider.Decrypt(ke, output, tmpEnc)

		// todo checksum, now pass directly

		utils.ArrayCopy(output, confoundedLen, tmpEnc, 0, dataLen)
		return tmpEnc[:dataLen]
	}else{  // reserved interface
		return nil
	}
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
