package cryptutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func HashSHA256(hashValue string, key string) (string, error) {
	h := hmac.New(sha256.New, []byte(key))

	_, err := h.Write([]byte(hashValue))
	if err != nil {
		return "", err
	}

	sha := hex.EncodeToString(h.Sum(nil))

	return sha, nil

}

func HashSHA512(hashValue string, key string) (string, error) {
	h := hmac.New(sha512.New, []byte(key))

	_, err := h.Write([]byte(hashValue))
	if err != nil {
		return "", err
	}

	sha := hex.EncodeToString(h.Sum(nil))

	return sha, nil

}
