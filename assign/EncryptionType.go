package assign

type EncryptionType struct {
	Value       int
	Name        string
	Description string
}

func NewEncryptionType(value int, name string, description string) *EncryptionType {
	return &EncryptionType{
		Value:       value,
		Name:        name,
		Description: description,
	}
}

func (this *EncryptionType) GetValue() int {
	return this.Value
}

func InitEnumType() []*EncryptionType {
	earray := make([]*EncryptionType, 33)
	earray[0] = NewEncryptionType(0, "none", "None encryption type")
	earray[1] = NewEncryptionType(1, "des-cbc-crc", "DES cbc mode with CRC-32")
	earray[2] = NewEncryptionType(2, "des-cbc-md4", "DES cbc mode with RSA-MD4")
	earray[3] = NewEncryptionType(3, "des-cbc-md5", "DES cbc mode with RSA-MD5")
	earray[4] = NewEncryptionType(3, "des", "DES cbc mode with RSA-MD5")
	earray[5] = NewEncryptionType(4, "des-cbc-raw", "DES cbc mode raw")
	earray[6] = NewEncryptionType(5, "des3-cbc-sha", "DES-3 cbc with SHA1")
	earray[7] = NewEncryptionType(6, "des3-cbc-raw", "Triple DES cbc mode raw")
	earray[8] = NewEncryptionType(8, "des-hmac-sha1", "DES with HMAC/sha1")
	earray[9] = NewEncryptionType(9, "dsa-sha1-cms", "DSA with SHA1, CMS signature")
	earray[10] = NewEncryptionType(10, "md5-rsa-cms", "MD5 with RSA, CMS signature")
	earray[11] = NewEncryptionType(11, "sha1-rsa-cms", "SHA1 with RSA, CMS signature")
	earray[12] = NewEncryptionType(12, "rc2-cbc-env", "RC2 cbc mode, CMS enveloped data")
	earray[13] = NewEncryptionType(13, "rsa-env", "RSA encryption, CMS enveloped data")
	earray[14] = NewEncryptionType(14, "rsa-es-oaep-env", "RSA w/OEAP encryption, CMS enveloped data")
	earray[15] = NewEncryptionType(15, "des3-cbc-env", "DES-3 cbc mode, CMS enveloped data")
	earray[16] = NewEncryptionType(16, "des3-cbc-sha1", "Triple DES cbc mode with HMAC/sha1")
	earray[17] = NewEncryptionType(16, "des3-hmac-sha1", "Triple DES cbc mode with HMAC/sha1")
	earray[18] = NewEncryptionType(16, "des3-cbc-sha1-kd", "Triple DES cbc mode with HMAC/sha1")
	earray[19] = NewEncryptionType(17, "aes128-cts-hmac-sha1-96", "AES-128 CTS mode with 96-bit SHA-1 HMAC")
	earray[20] = NewEncryptionType(17, "aes128-cts", "AES-128 CTS mode with 96-bit SHA-1 HMAC")
	earray[21] = NewEncryptionType(18, "aes256-cts-hmac-sha1-96", "AES-256 CTS mode with 96-bit SHA-1 HMAC")
	earray[22] = NewEncryptionType(18, "aes256-cts", "AES-256 CTS mode with 96-bit SHA-1 HMAC")
	earray[23] = NewEncryptionType(23, "arcfour-hmac", "ArcFour with HMAC/md5")
	earray[24] = NewEncryptionType(23, "rc4-hmac", "ArcFour with HMAC/md5")
	earray[25] = NewEncryptionType(23, "arcfour-hmac-md5", "ArcFour with HMAC/md5")
	earray[26] = NewEncryptionType(24, "arcfour-hmac-exp", "Exportable ArcFour with HMAC/md5")
	earray[27] = NewEncryptionType(24, "rc4-hmac-exp", "Exportable ArcFour with HMAC/md5")
	earray[28] = NewEncryptionType(24, "arcfour-hmac-md5-exp", "Exportable ArcFour with HMAC/md5")
	earray[29] = NewEncryptionType(25, "camellia128-cts-cmac", "Camellia-128 CTS mode with CMAC")
	earray[30] = NewEncryptionType(25, "camellia128-cts", "Camellia-128 CTS mode with CMAC")
	earray[31] = NewEncryptionType(26, "camellia256-cts-cmac", "Camellia-256 CTS mode with CMAC")
	earray[32] = NewEncryptionType(26, "camellia256-cts", "Camellia-256 CTS mode with CMAC")
	return earray
}

func EncryptionTypeFromValue(value int) *EncryptionType {
	erray := InitEnumType()
	for _, ele := range erray {
		if ele.GetValue() == value {
			return ele
		}
	}

	return &EncryptionType{Value:0, Name:"none", Description:"None encryption type"}
}
