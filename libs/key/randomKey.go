package genKey

import (
	"math/rand"
)

func RandomKey(keyLen int) string {
	var result = ""
	const letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const charLen int = len(letters)
	var counter = 0
	for counter < keyLen {
		result += string(rune(letters[rand.Intn(charLen)]))
		counter++
	}
	return result
}
