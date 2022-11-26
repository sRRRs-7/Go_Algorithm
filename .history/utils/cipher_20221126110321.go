package utils

import (
	"crypto/sha256"
	"fmt"
)

func Cipher() {
	res := HashPassword("srrrs")
	fmt.Println(res)
}

func HashPassword(password string) byte {
	sha := sha256.New()
	hash := sha.Sum([]byte(password))
	fmt.Println(hash)

	var s string
	for _, x := range hash {
		s += string(x)
	}
	return s
}