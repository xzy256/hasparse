package assign

type AbstractEncryptProvider interface {
	BlockSize() int //crypto block size
	ChecksumSize() int
	Decrypt(key, cipherState, data []byte)
	Encrypt(res []byte)
	KeyInputSize() int //input size to make key
	KeySize() int      //output key size
}
