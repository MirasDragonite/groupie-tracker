package main

import "groupie-tracker/internal/api"

func main() {
	api.Start()
}

func ReverseBits(oct byte) byte {
	var result byte

	for i := 0; i < 8; i++ {
		result <<= 1
		result |= oct & 1
		oct >>= 1
	}
	return result
}

func SwitchBits(oct byte) byte {
	leftHAlf := oct >> 4
	rightHalf := oct << 4
	return leftHAlf | rightHalf
}
