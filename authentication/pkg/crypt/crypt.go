package crypt

func Encrypt(text string) uint32 {

	var sum uint32 = 0

	for index, code := range text {
		sum += uint32((index + 1) * int(code))
	}
	return sum
}
