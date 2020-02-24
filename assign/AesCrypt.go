package assign

import (
	"github.com/xzy256/hasparse/utils"
	"log"
)

type AesCrypt struct {
	Alog      [256]int
	AesLog    [256]int
	K         []int
	T1        [256]int
	T2        [256]int
	T3        [256]int
	T4        [256]int
	T5        [256]int
	T6        [256]int
	T7        [256]int
	T8        [256]int
	U1        [256]int
	U2        [256]int
	U3        [256]int
	U4        [256]int
	S         [256]byte
	Si        [256]byte
	Rcon      [30]byte
	SessionK  [][]int
	ROUNDS_12 bool
	ROUNDS_14 bool
	Limit     int
	Dcpk      []byte
	Dcpr      []byte
}

func (this *AesCrypt) PreConstruct() {
	this.Dcpk = make([]byte, 16)
	this.Dcpr = make([]byte, 16)
	var0 := 283
	this.Alog[0] = 1
	var var1, var15 int
	for var1 = 1; var1 < 256; var1++ {
		var15 = this.Alog[var1-1]<<1 ^ this.Alog[var1-1]
		if (var15 & 256) != 0 {
			var15 ^= var0
		}

		this.Alog[var1] = var15
	}

	for var1 = 1; var1 < 255; {
		this.AesLog[this.Alog[var1]] = var1
		var1++
	}

	var3 := [][]byte{{1, 1, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 1, 1, 1, 0, 0}, {0, 0, 1, 1, 1, 1, 1, 0}, {0, 0, 0, 1, 1, 1, 1, 1}, {1, 0, 0, 0, 1, 1, 1, 1}, {1, 1, 0, 0, 0, 1, 1, 1}, {1, 1, 1, 0, 0, 0, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 1}}
	var4 := []byte{0, 1, 1, 0, 0, 0, 1, 1}
	var var6 [256][8]byte
	var6[1][7] = 1

	var var5 int
	for var1 = 2; var1 < 256; var1++ {
		var15 = this.Alog[255-this.AesLog[var1]]

		for var5 = 0; var5 < 8; var5++ {
			var6[var1][var5] = byte(uint(var15) >> (7 - var5) & 1)
		}
	}

	var var7 [256][8]byte

	for var1 = 0; var1 < 256; var1++ {
		for var5 = 0; var5 < 8; var5++ {
			var7[var1][var5] = var4[var5]

			for var15 = 0; var15 < 8; var15++ {
				var7[var1][var5] = var7[var1][var5] ^ var3[var5][var15]*var6[var1][var15]
			}
		}
	}

	for var1 = 0; var1 < 256; var1++ {
		this.S[var1] = var7[var1][0] << 7

		for var5 = 1; var5 < 8; var5++ {
			this.S[var1] = this.S[var1] ^ var7[var1][var5]<<(7-var5)
		}

		this.Si[this.S[var1]&255] = byte(var1)
	}

	var8 := [][]byte{{2, 1, 1, 3}, {3, 2, 1, 1}, {1, 3, 2, 1}, {1, 1, 3, 2}}
	var var9 [4][8]byte

	for var1 = 0; var1 < 4; var1++ {
		for var15 = 0; var15 < 4; var15++ {
			var9[var1][var15] = var8[var1][var15]
		}

		var9[var1][var1+4] = 1
	}

	var var12 [4][4]byte

	for var1 = 0; var1 < 4; var1++ {
		var10 := var9[var1][var1]
		if var10 == 0 {
			for var5 = var1 + 1; var9[var5][var1] == 0 && var5 < 4; var5++ {
			}

			if var5 == 4 {
				log.Fatal("G matrix is not invertible")
			}

			for var15 = 0; var15 < 8; var15++ {
				var11 := var9[var1][var15]
				var9[var1][var15] = var9[var5][var15]
				var9[var5][var15] = var11
			}

			var10 = var9[var1][var1]
		}

		for var15 = 0; var15 < 8; var15++ {
			if var9[var1][var15] != 0 {
				var9[var1][var15] = byte(this.Alog[(255+this.AesLog[var9[var1][var15]&255]-this.AesLog[var10&255])%255])
			}
		}

		for var5 = 0; var5 < 4; var5++ {
			if var1 != var5 {
				for var15 = var1 + 1; var15 < 8; var15++ {
					var9[var5][var15] = byte(int(var9[var5][var15]) ^ this.Mul(int(var9[var1][var15]), int(var9[var5][var1])))
				}

				var9[var5][var1] = 0
			}
		}
	}

	for var1 = 0; var1 < 4; var1++ {
		for var15 = 0; var15 < 4; var15++ {
			var12[var1][var15] = var9[var1][var15+4]
		}
	}

	for var5 = 0; var5 < 256; var5++ {
		var13 := int(this.S[var5])
		this.T1[var5] = this.Mul4(var13, var8[0])
		this.T2[var5] = this.Mul4(var13, var8[1])
		this.T3[var5] = this.Mul4(var13, var8[2])
		this.T4[var5] = this.Mul4(var13, var8[3])
		var13 = int(this.Si[var5])
		this.T5[var5] = this.Mul4(var13, var12[0][:])
		this.T6[var5] = this.Mul4(var13, var12[1][:])
		this.T7[var5] = this.Mul4(var13, var12[2][:])
		this.T8[var5] = this.Mul4(var13, var12[3][:])
		this.U1[var5] = this.Mul4(var5, var12[0][:])
		this.U2[var5] = this.Mul4(var5, var12[1][:])
		this.U3[var5] = this.Mul4(var5, var12[2][:])
		this.U4[var5] = this.Mul4(var5, var12[3][:])
	}

	this.Rcon[0] = 1
	var14 := 1
	for var5 = 1; var5 < 30; var5++ {
		var14 = this.Mul(2, var14)
		this.Rcon[var5] = byte(var14)
	}
}

func (this *AesCrypt) Mul4(var0 int, var1 []byte) int {
	if var0 == 0 {
		return 0
	} else {
		var0 = this.AesLog[var0&255]
		var var2, var3, var4, var5 int
		var2 = notZeroConf(var1[0], this.Alog[(var0+this.AesLog[var1[0]&255])%255]&255, 0)
		var3 = notZeroConf(var1[1], this.Alog[(var0+this.AesLog[var1[1]&255])%255]&255, 0)
		var4 = notZeroConf(var1[2], this.Alog[(var0+this.AesLog[var1[2]&255])%255]&255, 0)
		var5 = notZeroConf(var1[3], this.Alog[(var0+this.AesLog[var1[3]&255])%255]&255, 0)

		return var2<<24 | var3<<16 | var4<<8 | var5
	}
}

func notZeroConf(conf byte, val1 int, val2 int) int {
	if conf != 0 {
		return val1
	}
	return val2
}

func (this *AesCrypt) Mul(var0 int, var1 int) int {
	if var0 != 0 && var1 != 0 {
		return this.Alog[(this.AesLog[var0&255]+this.AesLog[var1&255])%255]
	}
	return 0
}

func isKeySizeValid(var0 int) bool {
	AES_KEYSIZES := []int{16, 24, 32}
	for var1 := 0; var1 < len(AES_KEYSIZES); var1++ {
		if var0 == AES_KEYSIZES[var1] {
			return true
		}
	}
	return false
}

func getRounds(var0 int) int {
	return (var0 >> 2) + 6
}

func expandToSubKey(var0 [][4]int, var1 bool) []int {
	var2 := len(var0)
	var3 := make([]int, var2*4)
	var var4, var5 int

	if var1 {
		for var4 = 0; var4 < 4; var4++ {
			var3[var4] = var0[var2-1][var4]
		}

		for var4 = 1; var4 < var2; var4++ {
			for var5 = 0; var5 < 4; var5++ {
				var3[var4*4+var5] = var0[var4-1][var5]
			}
		}
	} else {
		for var4 = 0; var4 < var2; var4++ {
			for var5 = 0; var5 < 4; var5++ {
				var3[var4*4+var5] = var0[var4][var5]
			}
		}
	}

	return var3
}

func (this *AesCrypt) MakeSessionKey(var1 []byte) {
	if var1 == nil {
		log.Fatal("Empty key")
	} else if !isKeySizeValid(len(var1)) {
		log.Fatal("Invalid AES key length: ", len(var1), " bytes")
	} else {
		var2 := getRounds(len(var1))
		var3 := (var2 + 1) * 4
		var var4 byte = 4
		var5 := make([][4]int, var2+1)
		var6 := make([][4]int, var2+1)
		var7 := len(var1) / 4
		var8 := make([]int, var7)
		var9 := 0

		var var10 int
		for var10 = 0; var9 < var7; var10 += 4 {
			var8[var9] = int(var1[var10])<<24 | int(var1[var10+1]&255)<<16 | int(var1[var10+2]&255)<<8 | int(var1[var10+3]&255)

			var9++
		}

		var11 := 0

		for var10 = 0; var10 < var7 && var11 < var3; var11++ {
			var5[var11/4][var11%4] = var8[var10]
			var6[var2-var11/4][var11%4] = var8[var10]
			var10++
		}

		var13 := 0

		var var12 int

		for var11 < var3 {
			var12 = var8[var7-1]
			t1 := uint(var12) >> 16 & 255
			t2 := uint(var12) >> 8 & 255

			op1 := int(this.S[t1]) << 24
			op2 := int(this.S[t2]&255) << 16
			op3 := int(this.S[var12&255]&255) << 8
			op4 := int(this.S[uint(var12)>>24]) & 255
			op5 := int(this.Rcon[var13]) << 24
			var8[0] ^= op1 ^ op2 ^ op3 ^ op4 ^ op5

			var13++
			if var7 != 8 {
				var9 = 1

				for var10 = 0; var9 < var7; var10++ {
					var8[var9] ^= var8[var10]
					var9++
				}
			} else {
				var9 = 1

				for var10 = 0; var9 < var7/2; var10++ {
					var8[var9] ^= var8[var10]
					var9++
				}

				var12 = var8[var7/2-1]

				var8[var7/2] ^= int(this.S[var12&255]&255 ^ (this.S[uint(var12)>>8&255]&255)<<8 ^ (this.S[uint(var12)>>16&255]&255)<<16 ^ this.S[uint(var12)>>24]<<24)
				var10 = var7 / 2

				for var9 = var10 + 1; var9 < var7; var10++ {
					var8[var9] ^= var8[var10]
					var9++
				}
			}

			for var10 = 0; var10 < var7 && var11 < var3; var11++ {
				var5[var11/4][var11%4] = var8[var10]
				var6[var2-var11/4][var11%4] = var8[var10]
				var10++
			}
		}

		for var14 := 1; var14 < var2; var14++ {
			for var10 = 0; var10 < int(var4); var10++ {
				var12 = var6[var14][var10]
				var6[var14][var10] = this.U1[uint(var12)>>24&255] ^ this.U2[uint(var12)>>16&255] ^ this.U3[uint(var12)>>8&255] ^ this.U4[var12&255]
			}
		}

		var16 := expandToSubKey(var5, false)
		var15 := expandToSubKey(var6, true)
		this.ROUNDS_12 = var2 >= 12
		this.ROUNDS_14 = var2 == 14
		this.Limit = var2 * 4
		this.SessionK = [][]int{var16, var15}
	}
}

func (this *AesCrypt) DecryptBlock2(var1 []byte, var2 int, var3 []byte, var4 int) {
	var var5 byte = 4
	var10000 := int(var1[var2])<<24 | int(var1[var2]&255)<<16 | int(var1[var2]&255)<<8 | int(var1[var2])&255
	var2 += 4
	var13 := var5 + 1
	var6 := var10000 ^ this.K[var5]
	var7 := (int(var1[var2])<<24 | int(var1[var2]&255)<<16 | int(var1[var2]&255)<<8 | int(var1[var2]&255)) ^ this.K[var13]
	var2 += 4
	var8 := int(var1[var2])<<24 | int(var1[var2]&255)<<16 | int(var1[var2]&255)<<8 | int(var1[var2]&255) ^ this.K[var13+1]
	var2 += 4
	var9 := int(var1[var2])<<24 | int(var1[var2]&255)<<16 | int(var1[var2]&255)<<8 | int(var1[var2]&255) ^ this.K[var13+2]
	var2 += 4
	var13 += 3
	var var10, var11, var12 int
	if this.ROUNDS_12 {
		var10 = this.T5[uint(var6)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var8)>>8&255] ^ this.T8[var7&255] ^ this.K[var13]
		var11 = this.T5[uint(var7)>>24] ^ this.T6[uint(var6)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var8&255] ^ this.K[var13+1]
		var12 = this.T5[uint(var8)>>24] ^ this.T6[uint(var7)>>16&255] ^ this.T7[uint(var6)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+2]
		var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var8)>>16&255] ^ this.T7[uint(var7)>>8&255] ^ this.T8[var6&255] ^ this.K[var13+3]
		var6 = this.T5[uint(var10)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var12)>>8&255] ^ this.T8[var11&255] ^ this.K[var13+4]
		var7 = this.T5[uint(var11)>>24] ^ this.T6[uint(var10)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var12&255] ^ this.K[var13+5]
		var8 = this.T5[uint(var12)>>24] ^ this.T6[uint(var11)>>16&255] ^ this.T7[uint(var10)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+6]
		var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var12)>>16&255] ^ this.T7[uint(var11)>>8&255] ^ this.T8[var10&255] ^ this.K[var13+7]
		var13 += 8
		if this.ROUNDS_14 {
			var10 = this.T5[uint(var6)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var8)>>8&255] ^ this.T8[var7&255] ^ this.K[var13]
			var11 = this.T5[uint(var7)>>24] ^ this.T6[uint(var6)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var8&255] ^ this.K[var13+1]
			var12 = this.T5[uint(var8)>>24] ^ this.T6[uint(var7)>>16&255] ^ this.T7[uint(var6)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+2]
			var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var8)>>16&255] ^ this.T7[uint(var7)>>8&255] ^ this.T8[var6&255] ^ this.K[var13+3]
			var6 = this.T5[uint(var10)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var12)>>8&255] ^ this.T8[var11&255] ^ this.K[var13+4]
			var7 = this.T5[uint(var11)>>24] ^ this.T6[uint(var10)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var12&255] ^ this.K[var13+5]
			var8 = this.T5[uint(var12)>>24] ^ this.T6[uint(var11)>>16&255] ^ this.T7[uint(var10)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+6]
			var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var12)>>16&255] ^ this.T7[uint(var11)>>8&255] ^ this.T8[var10&255] ^ this.K[var13+7]
			var13 += 8
		}
	}

	var10 = this.T5[uint(var6)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var8)>>8&255] ^ this.T8[var7&255] ^ this.K[var13]
	var11 = this.T5[uint(var7)>>24] ^ this.T6[uint(var6)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var8&255] ^ this.K[var13+1]
	var12 = this.T5[uint(var8)>>24] ^ this.T6[uint(var7)>>16&255] ^ this.T7[uint(var6)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+2]
	var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var8)>>16&255] ^ this.T7[uint(var7)>>8&255] ^ this.T8[var6&255] ^ this.K[var13+3]
	var6 = this.T5[uint(var10)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var12)>>8&255] ^ this.T8[var11&255] ^ this.K[var13+4]
	var7 = this.T5[uint(var11)>>24] ^ this.T6[uint(var10)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var12&255] ^ this.K[var13+5]
	var8 = this.T5[uint(var12)>>24] ^ this.T6[uint(var11)>>16&255] ^ this.T7[uint(var10)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+6]
	var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var12)>>16&255] ^ this.T7[uint(var11)>>8&255] ^ this.T8[var10&255] ^ this.K[var13+7]
	var10 = this.T5[uint(var6)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var8)>>8&255] ^ this.T8[var7&255] ^ this.K[var13+8]
	var11 = this.T5[uint(var7)>>24] ^ this.T6[uint(var6)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var8&255] ^ this.K[var13+9]
	var12 = this.T5[uint(var8)>>24] ^ this.T6[uint(var7)>>16&255] ^ this.T7[uint(var6)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+10]
	var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var8)>>16&255] ^ this.T7[uint(var7)>>8&255] ^ this.T8[var6&255] ^ this.K[var13+11]
	var6 = this.T5[uint(var10)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var12)>>8&255] ^ this.T8[var11&255] ^ this.K[var13+12]
	var7 = this.T5[uint(var11)>>24] ^ this.T6[uint(var10)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var12&255] ^ this.K[var13+13]
	var8 = this.T5[uint(var12)>>24] ^ this.T6[uint(var11)>>16&255] ^ this.T7[uint(var10)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+14]
	var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var12)>>16&255] ^ this.T7[uint(var11)>>8&255] ^ this.T8[var10&255] ^ this.K[var13+15]
	var10 = this.T5[uint(var6)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var8)>>8&255] ^ this.T8[var7&255] ^ this.K[var13+16]
	var11 = this.T5[uint(var7)>>24] ^ this.T6[uint(var6)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var8&255] ^ this.K[var13+17]
	var12 = this.T5[uint(var8)>>24] ^ this.T6[uint(var7)>>16&255] ^ this.T7[uint(var6)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+18]
	var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var8)>>16&255] ^ this.T7[uint(var7)>>8&255] ^ this.T8[var6&255] ^ this.K[var13+19]
	var6 = this.T5[uint(var10)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var12)>>8&255] ^ this.T8[var11&255] ^ this.K[var13+20]
	var7 = this.T5[uint(var11)>>24] ^ this.T6[uint(var10)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var12&255] ^ this.K[var13+21]
	var8 = this.T5[uint(var12)>>24] ^ this.T6[uint(var11)>>16&255] ^ this.T7[uint(var10)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+22]
	var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var12)>>16&255] ^ this.T7[uint(var11)>>8&255] ^ this.T8[var10&255] ^ this.K[var13+23]
	var10 = this.T5[uint(var6)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var8)>>8&255] ^ this.T8[var7&255] ^ this.K[var13+24]
	var11 = this.T5[uint(var7)>>24] ^ this.T6[uint(var6)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var8&255] ^ this.K[var13+25]
	var12 = this.T5[uint(var8)>>24] ^ this.T6[uint(var7)>>16&255] ^ this.T7[uint(var6)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+26]
	var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var8)>>16&255] ^ this.T7[uint(var7)>>8&255] ^ this.T8[var6&255] ^ this.K[var13+27]
	var6 = this.T5[uint(var10)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var12)>>8&255] ^ this.T8[var11&255] ^ this.K[var13+28]
	var7 = this.T5[uint(var11)>>24] ^ this.T6[uint(var10)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var12&255] ^ this.K[var13+29]
	var8 = this.T5[uint(var12)>>24] ^ this.T6[uint(var11)>>16&255] ^ this.T7[uint(var10)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+30]
	var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var12)>>16&255] ^ this.T7[uint(var11)>>8&255] ^ this.T8[var10&255] ^ this.K[var13+31]
	var10 = this.T5[uint(var6)>>24] ^ this.T6[uint(var9)>>16&255] ^ this.T7[uint(var8)>>8&255] ^ this.T8[var7&255] ^ this.K[var13+32]
	var11 = this.T5[uint(var7)>>24] ^ this.T6[uint(var6)>>16&255] ^ this.T7[uint(var9)>>8&255] ^ this.T8[var8&255] ^ this.K[var13+33]
	var12 = this.T5[uint(var8)>>24] ^ this.T6[uint(var7)>>16&255] ^ this.T7[uint(var6)>>8&255] ^ this.T8[var9&255] ^ this.K[var13+34]
	var9 = this.T5[uint(var9)>>24] ^ this.T6[uint(var8)>>16&255] ^ this.T7[uint(var7)>>8&255] ^ this.T8[var6&255] ^ this.K[var13+35]
	var13 += 36
	var7 = this.K[0]
	var3[var4] = (byte)(uint(this.Si[uint(var10)>>24]) ^ uint(var7)>>24)
	var3[var4+1] = (byte)(uint(this.Si[uint(var9)>>16&255]) ^ uint(var7)>>16)
	var3[var4+2] = (byte)(uint(this.Si[uint(var12)>>8&255]) ^ uint(var7)>>8)
	var3[var4+3] = (byte)(int(this.Si[var11&255]) ^ var7)
	var7 = this.K[1]
	var3[var4+4] = (byte)(uint(this.Si[uint(var11)>>24]) ^ uint(var7)>>24)
	var3[var4+5] = (byte)(uint(this.Si[uint(var10)>>16&255]) ^ uint(var7)>>16)
	var3[var4+6] = (byte)(uint(this.Si[uint(var9)>>8&255]) ^ uint(var7)>>8)
	var3[var4+7] = (byte)(int(this.Si[var12&255]) ^ var7)
	var7 = this.K[2]
	var3[var4+8] = (byte)(uint(this.Si[uint(var12)>>24]) ^ uint(var7)>>24)
	var3[var4+9] = (byte)(uint(this.Si[uint(var11)>>16&255]) ^ uint(var7)>>16)
	var3[var4+10] = (byte)(uint(this.Si[uint(var10)>>8&255]) ^ uint(var7)>>8)
	var3[var4+11] = (byte)(int(this.Si[var9&255]) ^ var7)
	var7 = this.K[3]
	var3[var4+12] = (byte)(uint(this.Si[uint(var9)>>24]) ^ uint(var7)>>24)
	var3[var4+13] = (byte)(uint(this.Si[uint(var12)>>16&255]) ^ uint(var7)>>16)
	var3[var4+14] = (byte)(uint(this.Si[uint(var11)>>8&255]) ^ uint(var7)>>8)
	var4 += 15
	var3[var4] = (byte)(int(this.Si[var10&255]) ^ var7)
}

func (this *AesCrypt) EncryptBlock(paramArrayOfByte1 []byte, paramInt1 int, paramArrayOfByte2 []byte, paramInt2 int) {
	i := 0
	j := (int(paramArrayOfByte1[paramInt1])<<24 | int(paramArrayOfByte1[paramInt1+1]&0xFF)<<16 | int(paramArrayOfByte1[paramInt1+2]&0xFF)<<8 | int(paramArrayOfByte1[paramInt1+3]&0xFF)) ^ this.K[i]
	k := (int(paramArrayOfByte1[paramInt1+4])<<24 | int(paramArrayOfByte1[paramInt1+5]&0xFF)<<16 | int(paramArrayOfByte1[paramInt1+6]&0xFF)<<8 | int(paramArrayOfByte1[paramInt1+7]&0xFF)) ^ this.K[i+1]
	m := (int(paramArrayOfByte1[paramInt1+8])<<24 | int(paramArrayOfByte1[paramInt1+9]&0xFF)<<16 | int(paramArrayOfByte1[paramInt1+10]&0xFF)<<8 | int(paramArrayOfByte1[paramInt1+11]&0xFF)) ^ this.K[i+2]
	n := (int(paramArrayOfByte1[paramInt1+12])<<24 | int(paramArrayOfByte1[paramInt1+13]&0xFF)<<16 | int(paramArrayOfByte1[paramInt1+14]&0xFF)<<8 | int(paramArrayOfByte1[paramInt1+15]&0xFF)) ^ this.K[i+3]
	i += 4
	paramInt1 += 16
	var i3 int
	for ; i < this.Limit; m = i3 {
		i1 := this.T1[uint(j)>>24] ^ this.T2[uint(k)>>16&0xFF] ^ this.T3[uint(m)>>8&0xFF] ^ this.T4[n&0xFF] ^ this.K[i]
		i2 := this.T1[uint(k)>>24] ^ this.T2[uint(m)>>16&0xFF] ^ this.T3[uint(n)>>8&0xFF] ^ this.T4[j&0xFF] ^ this.K[i+1]
		i3 = this.T1[uint(m)>>24] ^ this.T2[uint(n)>>16&0xFF] ^ this.T3[uint(j)>>8&0xFF] ^ this.T4[k&0xFF] ^ this.K[i+2]
		n = this.T1[uint(n)>>24] ^ this.T2[uint(j)>>16&0xFF] ^ this.T3[uint(k)>>8&0xFF] ^ this.T4[m&0xFF] ^ this.K[i+3]
		i += 4
		j = i1
		k = i2
	}
	i1 := this.K[i]
	paramArrayOfByte2[paramInt2] = byte(uint(this.S[uint(j)>>24]) ^ uint(i1)>>24)
	paramArrayOfByte2[paramInt2+1] = byte(uint(this.S[uint(k)>>16&0xFF]) ^ uint(i1)>>16)
	paramArrayOfByte2[paramInt2+2] = byte(uint(this.S[uint(m)>>8&0xFF]) ^ uint(i1)>>8)
	paramArrayOfByte2[paramInt2+3] = byte(int(this.S[(n&0xFF)]) ^ i1)
	i1 = this.K[i+1]
	paramArrayOfByte2[paramInt2+4] = byte(uint(this.S[uint(k)>>24]) ^ uint(i1)>>24)
	paramArrayOfByte2[paramInt2+5] = byte(uint(this.S[uint(m)>>16&0xFF]) ^ uint(i1)>>16)
	paramArrayOfByte2[paramInt2+6] = byte(uint(this.S[uint(n)>>8&0xFF]) ^ uint(i1)>>8)
	paramArrayOfByte2[paramInt2+7] = byte(int(this.S[j&0xFF]) ^ i1)
	i1 = this.K[i+2]
	paramArrayOfByte2[paramInt2+8] = byte(uint(this.S[uint(m)>>24]) ^ uint(i1)>>24)
	paramArrayOfByte2[paramInt2+9] = byte(uint(this.S[uint(n)>>16&0xFF]) ^ uint(i1)>>16)
	paramArrayOfByte2[paramInt2+10] = byte(uint(this.S[uint(j)>>8&0xFF]) ^ uint(i1)>>8)
	paramArrayOfByte2[paramInt2+11] = byte(int(this.S[k&0xFF]) ^ i1)
	i1 = this.K[i+3]
	paramArrayOfByte2[paramInt2+12] = byte(uint(this.S[uint(n)>>24]) ^ uint(i1)>>24)
	paramArrayOfByte2[paramInt2+13] = byte(uint(this.S[uint(j)>>16&0xFF]) ^ uint(i1)>>16)
	paramArrayOfByte2[paramInt2+14] = byte(uint(this.S[uint(k)>>8&0xFF]) ^ uint(i1)>>8)
	paramInt2 += 15
	i += 4
	paramArrayOfByte2[paramInt2] = byte(int(this.S[m&0xFF]) ^ i1)
}

func (this *AesCrypt) DecryptFinal(var1 []byte, var2 int, var3 int, var4 []byte, var5 int, blockSize int) {
	if var3 < blockSize {
		log.Fatal("Input is too short!")
	} else if var3 == blockSize {
		this.decrypt(var1, var2, var3, var4, var5, blockSize)
	} else {
		var6 := var3 % blockSize
		var var7 int
		if var6 == 0 { // uncheck
			var7 = var2 + var3 - blockSize
			var8 := var2 + var3 - 2*blockSize
			var9 := make([]byte, 2*blockSize)
			utils.ArrayCopy(var1, var7, var9, 0, blockSize)
			utils.ArrayCopy(var1, var8, var9, blockSize, blockSize)
			var10 := var3 - 2*blockSize
			this.decrypt(var1, var2, var10, var4, var5, blockSize)
			this.decrypt(var9, 0, 2*blockSize, var4, var5+var10, blockSize)
		} else {
			var7 = var3 - (blockSize + var6)
			if var7 > 0 {
				this.decrypt(var1, var2, var7, var4, var5, blockSize)
				var2 += var7
				var5 += var7
			}
			var11 := make([]byte, blockSize)
			this.DecryptBlock(var1, var2, var11, 0)
			var12 := 0
			for var12 = 0; var12 < var6; var12++ {
				var4[var5+blockSize+var12] = var1[var2+blockSize+var12] ^ var11[var12]
			}
			utils.ArrayCopy(var1, var2+blockSize, var11, 0, var6)
			this.DecryptBlock(var11, 0, var4, var5)
			for var12 = 0; var12 < blockSize; var12++ {
				var4[var5+var12] ^= this.Dcpr[var12]
			}
		}
	}
}

func (this *AesCrypt) decrypt(var1 []byte, var2 int, var3 int, var4 []byte, var5 int, blockSize int) {
	if var3 <= 0 {
		return
	} else if var3%blockSize != 0 {
		log.Fatal("Internal error in input buffering")
	} else {
		for var6 := var2 + var3; var2 < var6; var5 += blockSize {
			this.DecryptBlock(var1, var2, this.Dcpk, 0)
			for var7 := 0; var7 < blockSize; var7++ {
				var4[var7+var5] = this.Dcpk[var7] ^ this.Dcpr[var7]
			}

			utils.ArrayCopy(var1, var2, this.Dcpr, 0, blockSize)
			var2 += blockSize
		}
	}
}

func (this *AesCrypt) DecryptBlock(paramArrayOfByte1 []byte, paramInt1 int, paramArrayOfByte2 []byte, paramInt2 int) {
	i := 4
	j := (int(paramArrayOfByte1[paramInt1])<<24 | int(paramArrayOfByte1[paramInt1+1]&0xFF)<<16 | int(paramArrayOfByte1[paramInt1+2]&0xFF)<<8 | int(paramArrayOfByte1[paramInt1+3]&0xFF)) ^ this.K[i]
	k := (int(paramArrayOfByte1[paramInt1+4])<<24 | int(paramArrayOfByte1[paramInt1+5]&0xFF)<<16 | int(paramArrayOfByte1[paramInt1+6]&0xFF)<<8 | int(paramArrayOfByte1[paramInt1+7]&0xFF)) ^ this.K[i+1]
	m := (int(paramArrayOfByte1[paramInt1+8])<<24 | int(paramArrayOfByte1[paramInt1+9]&0xFF)<<16 | int(paramArrayOfByte1[paramInt1+10]&0xFF)<<8 | int(paramArrayOfByte1[paramInt1+11]&0xFF)) ^ this.K[i+2]
	n := (int(paramArrayOfByte1[paramInt1+12])<<24 | int(paramArrayOfByte1[paramInt1+13]&0xFF)<<16 | int(paramArrayOfByte1[paramInt1+14]&0xFF)<<8 | int(paramArrayOfByte1[paramInt1+15]&0xFF)) ^ this.K[i+3]
	i += 4
	paramInt1 += 16

	var i1, i2, i3 int
	if this.ROUNDS_12 { // uncheck
		i1 = this.T5[uint(j)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(m)>>8&0xFF] ^ this.T8[k&0xFF] ^ this.K[i]
		i2 = this.T5[uint(k)>>24] ^ this.T6[uint(j)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[m&0xFF] ^ this.K[i+1]
		i3 = this.T5[uint(m)>>24] ^ this.T6[uint(k)>>16&0xFF] ^ this.T7[uint(j)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+2]
		n = this.T5[uint(n)>>24] ^ this.T6[uint(m)>>16&0xFF] ^ this.T7[uint(k)>>8&0xFF] ^ this.T8[j&0xFF] ^ this.K[i+3]
		j = this.T5[uint(i1)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(i3)>>8&0xFF] ^ this.T8[i2&0xFF] ^ this.K[i+4]
		k = this.T5[uint(i2)>>24] ^ this.T6[uint(i1)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[i3&0xFF] ^ this.K[i+5]
		m = this.T5[uint(i3)>>24] ^ this.T6[uint(i2)>>16&0xFF] ^ this.T7[uint(i1)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+6]
		n = this.T5[uint(n)>>24] ^ this.T6[uint(i3)>>16&0xFF] ^ this.T7[uint(i2)>>8&0xFF] ^ this.T8[i1&0xFF] ^ this.K[i+7]
		i += 8
		if this.ROUNDS_14 {
			i1 = this.T5[uint(j)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(m)>>8&0xFF] ^ this.T8[k&0xFF] ^ this.K[i]
			i2 = this.T5[uint(k)>>24] ^ this.T6[uint(j)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[m&0xFF] ^ this.K[i+1]
			i3 = this.T5[uint(m)>>24] ^ this.T6[uint(k)>>16&0xFF] ^ this.T7[uint(j)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+2]
			n = this.T5[uint(n)>>24] ^ this.T6[uint(m)>>16&0xFF] ^ this.T7[uint(k)>>8&0xFF] ^ this.T8[j&0xFF] ^ this.K[i+3]
			j = this.T5[uint(i1)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(i3)>>8&0xFF] ^ this.T8[i2&0xFF] ^ this.K[i+4]
			k = this.T5[uint(i2)>>24] ^ this.T6[uint(i1)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[i3&0xFF] ^ this.K[i+5]
			m = this.T5[uint(i3)>>24] ^ this.T6[uint(i2)>>16&0xFF] ^ this.T7[uint(i1)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+6]
			n = this.T5[uint(n)>>24] ^ this.T6[uint(i3)>>16&0xFF] ^ this.T7[uint(i2)>>8&0xFF] ^ this.T8[i1&0xFF] ^ this.K[i+7]
			i += 8
		}
	}
	i1 = this.T5[uint(j)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(m)>>8&0xFF] ^ this.T8[uint(k)&0xFF] ^ this.K[i]
	i2 = this.T5[uint(k)>>24] ^ this.T6[uint(j)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[uint(m)&0xFF] ^ this.K[i+1]
	i3 = this.T5[uint(m)>>24] ^ this.T6[uint(k)>>16&0xFF] ^ this.T7[uint(j)>>8&0xFF] ^ this.T8[uint(n)&0xFF] ^ this.K[i+2]
	n = this.T5[uint(n)>>24] ^ this.T6[uint(m)>>16&0xFF] ^ this.T7[uint(k)>>8&0xFF] ^ this.T8[uint(j)&0xFF] ^ this.K[i+3]
	j = this.T5[uint(i1)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(i3)>>8&0xFF] ^ this.T8[uint(i2)&0xFF] ^ this.K[i+4]
	k = this.T5[uint(i2)>>24] ^ this.T6[uint(i1)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[uint(i3)&0xFF] ^ this.K[i+5]
	m = this.T5[uint(i3)>>24] ^ this.T6[uint(i2)>>16&0xFF] ^ this.T7[uint(i1)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+6]
	n = this.T5[uint(n)>>24] ^ this.T6[uint(i3)>>16&0xFF] ^ this.T7[uint(i2)>>8&0xFF] ^ this.T8[i1&0xFF] ^ this.K[i+7]
	i1 = this.T5[uint(j)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(m)>>8&0xFF] ^ this.T8[k&0xFF] ^ this.K[i+8]

	i2 = this.T5[uint(k)>>24] ^ this.T6[uint(j)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[m&0xFF] ^ this.K[i+9]
	i3 = this.T5[uint(m)>>24] ^ this.T6[uint(k)>>16&0xFF] ^ this.T7[uint(j)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+10]
	n = this.T5[uint(n)>>24] ^ this.T6[uint(m)>>16&0xFF] ^ this.T7[uint(k)>>8&0xFF] ^ this.T8[j&0xFF] ^ this.K[i+11]
	j = this.T5[uint(i1)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(i3)>>8&0xFF] ^ this.T8[i2&0xFF] ^ this.K[i+12]
	k = this.T5[uint(i2)>>24] ^ this.T6[uint(i1)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[i3&0xFF] ^ this.K[i+13]
	m = this.T5[uint(i3)>>24] ^ this.T6[uint(i2)>>16&0xFF] ^ this.T7[uint(i1)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+14]
	n = this.T5[uint(n)>>24] ^ this.T6[uint(i3)>>16&0xFF] ^ this.T7[uint(i2)>>8&0xFF] ^ this.T8[i1&0xFF] ^ this.K[i+15]
	i1 = this.T5[uint(j)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(m)>>8&0xFF] ^ this.T8[k&0xFF] ^ this.K[i+16]
	i2 = this.T5[uint(k)>>24] ^ this.T6[uint(j)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[m&0xFF] ^ this.K[i+17]

	i3 = this.T5[uint(m)>>24] ^ this.T6[uint(k)>>16&0xFF] ^ this.T7[uint(j)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+18]

	n = this.T5[uint(n)>>24] ^ this.T6[uint(m)>>16&0xFF] ^ this.T7[uint(k)>>8&0xFF] ^ this.T8[j&0xFF] ^ this.K[i+19]

	j = this.T5[uint(i1)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(i3)>>8&0xFF] ^ this.T8[i2&0xFF] ^ this.K[i+20]

	k = this.T5[uint(i2)>>24] ^ this.T6[uint(i1)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[i3&0xFF] ^ this.K[i+21]

	m = this.T5[uint(i3)>>24] ^ this.T6[uint(i2)>>16&0xFF] ^ this.T7[uint(i1)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+22]

	n = this.T5[uint(n)>>24] ^ this.T6[uint(i3)>>16&0xFF] ^ this.T7[uint(i2)>>8&0xFF] ^ this.T8[i1&0xFF] ^ this.K[i+23]

	i1 = this.T5[uint(j)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(m)>>8&0xFF] ^ this.T8[k&0xFF] ^ this.K[i+24]

	i2 = this.T5[uint(k)>>24] ^ this.T6[uint(j)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[m&0xFF] ^ this.K[i+25]

	i3 = this.T5[uint(m)>>24] ^ this.T6[uint(k)>>16&0xFF] ^ this.T7[uint(j)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+26]

	n = this.T5[uint(n)>>24] ^ this.T6[uint(m)>>16&0xFF] ^ this.T7[uint(k)>>8&0xFF] ^ this.T8[j&0xFF] ^ this.K[i+27]

	j = this.T5[uint(i1)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(i3)>>8&0xFF] ^ this.T8[i2&0xFF] ^ this.K[i+28]

	k = this.T5[uint(i2)>>24] ^ this.T6[uint(i1)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[i3&0xFF] ^ this.K[i+29]

	m = this.T5[uint(i3)>>24] ^ this.T6[uint(i2)>>16&0xFF] ^ this.T7[uint(i1)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+30]

	n = this.T5[uint(n)>>24] ^ this.T6[uint(i3)>>16&0xFF] ^ this.T7[uint(i2)>>8&0xFF] ^ this.T8[i1&0xFF] ^ this.K[i+31]

	i1 = this.T5[uint(j)>>24] ^ this.T6[uint(n)>>16&0xFF] ^ this.T7[uint(m)>>8&0xFF] ^ this.T8[k&0xFF] ^ this.K[i+32]

	i2 = this.T5[uint(k)>>24] ^ this.T6[uint(j)>>16&0xFF] ^ this.T7[uint(n)>>8&0xFF] ^ this.T8[m&0xFF] ^ this.K[i+33]

	i3 = this.T5[uint(m)>>24] ^ this.T6[uint(k)>>16&0xFF] ^ this.T7[uint(j)>>8&0xFF] ^ this.T8[n&0xFF] ^ this.K[i+34]

	n = this.T5[uint(n)>>24] ^ this.T6[uint(m)>>16&0xFF] ^ this.T7[uint(k)>>8&0xFF] ^ this.T8[j&0xFF] ^ this.K[i+35]
	i += 36
	k = this.K[0]
	paramArrayOfByte2[paramInt2] = byte(uint(this.Si[uint(i1)>>24]) ^ uint(k)>>24)
	paramArrayOfByte2[paramInt2+1] = byte(uint(this.Si[uint(n)>>16&0xFF]) ^ uint(k)>>16)
	paramArrayOfByte2[paramInt2+2] = byte(uint(this.Si[uint(i3)>>8&0xFF]) ^ uint(k)>>8)
	paramArrayOfByte2[paramInt2+3] = byte(int(this.Si[i2&0xFF]) ^ k)
	paramInt2 += 4
	k = this.K[1]
	paramArrayOfByte2[paramInt2] = byte(uint(this.Si[uint(i2)>>24]) ^ uint(k)>>24)
	paramArrayOfByte2[paramInt2+1] = byte(uint(this.Si[uint(i1)>>16&0xFF]) ^ uint(k)>>16)
	paramArrayOfByte2[paramInt2+2] = byte(uint(this.Si[uint(n)>>8&0xFF]) ^ uint(k)>>8)
	paramArrayOfByte2[paramInt2+3] = byte(int(this.Si[i3&0xFF]) ^ k)
	paramInt2 += 4
	k = this.K[2]
	paramArrayOfByte2[paramInt2] = byte(uint(this.Si[uint(i3)>>24]) ^ uint(k)>>24)
	paramArrayOfByte2[paramInt2+1] = byte(uint(this.Si[uint(i2)>>16&0xFF]) ^ uint(k)>>16)
	paramArrayOfByte2[paramInt2+2] = byte(uint(this.Si[uint(i1)>>8&0xFF]) ^ uint(k)>>8)
	paramArrayOfByte2[paramInt2+3] = byte(int(this.Si[n&0xFF]) ^ k)
	paramInt2 += 4
	k = this.K[3]
	paramArrayOfByte2[paramInt2] = byte(uint(this.Si[uint(n)>>24]) ^ uint(k)>>24)
	paramArrayOfByte2[paramInt2+1] = byte(uint(this.Si[uint(i3)>>16&0xFF]) ^ uint(k)>>16)
	paramArrayOfByte2[paramInt2+2] = byte(uint(this.Si[uint(i2)>>8&0xFF]) ^ uint(k)>>8)
	paramArrayOfByte2[paramInt2+3] = byte(int(this.Si[i1&0xFF]) ^ k)
}
