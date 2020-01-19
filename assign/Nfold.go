package assign

/*
* representation: msb first, assume n and k are multiples of 8, and
* that {@code k>=16}.  this is the case of all the crypto systems which are
* likely to be used.  this function can be replaced if that
* assumption ever fails.
*/

func Selfnfold(inBytes []byte, size int) []byte {
	inBytesNum := len(inBytes) // count inBytes byte
	outBytesNum := size        // count outBytes byte

	a := outBytesNum
	b := inBytesNum

	for b != 0 {
		c := b
		b = a % b
		a = c
	}
	lcm := (outBytesNum * inBytesNum) / a

	outBytes := make([]byte, outBytesNum)
	for i := range outBytes { // as if ignore
		outBytes[i] = byte(0)
	}

	tmpByte := 0
	for i := lcm - 1; i >= 0; i-- {
		// first, start with the msbit inBytes the first, unrotated byte
		tmp := (inBytesNum << 3) - 1
		// then, for each byte, shift to the right for each repetition
		tmp += ((inBytesNum << 3) + 13) * (i / inBytesNum)
		// last, pick outBytes the correct byte within that shifted repetition
		tmp += (inBytesNum - (i % inBytesNum)) << 3

		msbit := tmp % (inBytesNum << 3)
		tmpmsbit := int(uint32(msbit) >> 3)
		// pull outBytes the byte value itself
		leftv := (int(inBytes[((inBytesNum-1)-tmpmsbit)%inBytesNum]) & 0xff) << 8
		rightv := int(inBytes[((inBytesNum)-tmpmsbit)%inBytesNum] & 0xff)
		transbits := uint32(msbit&7) + 1
		tmp = int(uint32(leftv|rightv)>>transbits) & 0xff

		tmpByte += tmp
		tmp = int(outBytes[i%outBytesNum]) & 0xff
		tmpByte += tmp
		outBytes[i%outBytesNum] = (byte)(tmpByte & 0xff)
		tmpByte = int(uint32(tmpByte) >> 8)
	}

	// if there's a carry bit left over, add it back inBytes
	if tmpByte != 0 {
		for i := outBytesNum - 1; i >= 0; i-- {
			// do the addition
			tmpByte += int(outBytes[i]) & 0xff
			outBytes[i] = (byte)(tmpByte & 0xff)

			tmpByte = int(uint32(tmpByte) >> 8)
		}
	}

	return outBytes
}
