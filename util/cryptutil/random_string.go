package cryptutil

import (
	"math/rand"
	"time"
)

const (
	RandomAlphanumType = "alphanum"
	RandomAlphaType    = "alpha"
	RandomNumberType   = "number"
)

func RandomString(strSize int, randType string) string {
	var dictionary string

	switch randType {
	case RandomAlphanumType:
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case RandomAlphaType:
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case RandomNumberType:
		dictionary = "0123456789"
	default:
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	var bytes = make([]byte, strSize)
	rand.Seed(time.Now().UnixNano())
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}
