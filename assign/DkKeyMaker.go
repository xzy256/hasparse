package assign

import (
	"github.com/xzy256/hasparse/utils"
)

/* Reference DkKeyMaker.java
 * K1 = E(Key, n-fold(Constant))
 * K2 = E(Key, K1)
 * K3 = E(Key, K2)
 * K4 = ...
 * DR(Key, Constant) = k-truncate(K1 | K2 | K3 | K4 ...)
 */
func Dr(key []byte, constant []byte, encProvider *Aes128Provider) []byte {
	encProvider.Aes128Crypt.MakeSessionKey(key)

	blocksize := encProvider.BlockSize()
	keyInuptSize := encProvider.KeyInputSize()
	keyBytes := make([]byte, keyInuptSize)
	var ki []byte

	if len(constant) != blocksize {
		ki = Selfnfold(constant, blocksize)
	} else {
		utils.ArrayCopy(constant, 0, ki, 0, len(constant))
	}

	n := 0
	for n < keyInuptSize {
		encProvider.Encrypt(ki)

		if n+blocksize >= keyInuptSize {
			utils.ArrayCopy(ki, 0, keyBytes, n, keyInuptSize-n)
			break
		}

		utils.ArrayCopy(ki, 0, keyBytes, n, blocksize)
		n += blocksize
	}

	return keyBytes
}
