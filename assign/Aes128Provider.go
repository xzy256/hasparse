package assign

import "hasparse/utils"

type Aes128Provider struct {
	Aes128Crypt *AesCrypt
}

func (this *Aes128Provider) Construct() {
	this.Aes128Crypt = &AesCrypt{}
	this.Aes128Crypt.PreConstruct()
}

func (this *Aes128Provider) KeyInputSize() int {
	return 16
}

func (this *Aes128Provider) KeySize() int {
	return 16
}

func (this *Aes128Provider) BlockSize() int {
	return 16
}

func (this *Aes128Provider) ChecksumSize() int {
	return 96 / 8
}

func (this *Aes128Provider) Encrypt(res []byte) {
	this.Aes128Crypt.K = this.Aes128Crypt.SessionK[0]
	output := make([]byte, 16)
	this.Aes128Crypt.EncryptBlock(res, 0, output, 0)
	utils.ArrayCopy(output, 0, res, 0, len(output))
}

func (this *Aes128Provider) Decrypt(key, output, data []byte) {
	this.Aes128Crypt.MakeSessionKey(key)
	this.Aes128Crypt.K = this.Aes128Crypt.SessionK[1]
	this.Aes128Crypt.DecryptFinal(data, 0, len(data), output, 0, this.BlockSize())
}
