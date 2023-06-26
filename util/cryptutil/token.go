package cryptutil

import (
	"encoding/base64"
)

type HashToken struct {
	Token string
	Hash  string
}

func IssueHashToken() (*HashToken, error) {
	token:= RandToken(128)
	hashStr, err := HashPassword(token,10)
	if err != nil{
		return nil, err
	}
	//encode to base64
	hashEncodeBase64 := base64.StdEncoding.EncodeToString([]byte(hashStr))
	hashToken := &HashToken{
		Token:token,
		Hash:hashEncodeBase64,
	}
	return hashToken, nil
}

func ValidHashToken(hashToken *HashToken) bool{
	if hashToken == nil{
		return false
	}

	hashDecodeByte, err := base64.StdEncoding.DecodeString(hashToken.Hash)
	if err != nil{
		return false
	}

	return VerifyHashPassword(hashToken.Token, string(hashDecodeByte))
}
