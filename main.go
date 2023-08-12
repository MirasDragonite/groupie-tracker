package main

import "groupie-tracker/cmd/api"

func main() {
	api.Start()
}

func SwapBits(oct byte) byte {
	leftHalf := oct >> 4
	rightHalf := oct >> 4
	return leftHalf | rightHalf
}

func ReverseBitd(oct byte) byte {
	var result byte

	for i := 0; i < 8; i++ {

		result <<= 1
		result |= oct & 1
		oct >>= 1
	}
	return result
}
