package rsautil

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/ssh"
)

const (
	rsaKeySize = 2048
)

func hash(data []byte) []byte {
	s := sha1.Sum(data)
	return s[:]

}

func GenerateSSHKeyPair() ([]byte, []byte, error) {
	privateKey, publicKey, err := generateKey()
	if err != nil {
		return nil, nil, err
	}
	privatePEM := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	privateBuff := &bytes.Buffer{}
	if err := pem.Encode(privateBuff, privatePEM); err != nil {
		return nil, nil, err
	}

	sshPub, err := ssh.NewPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}

	publicBytes := ssh.MarshalAuthorizedKey(sshPub)
	return privateBuff.Bytes(), publicBytes, nil

}

func generateKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	pri, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return nil, nil, err
	}
	return pri, &pri.PublicKey, nil
}

func GenerateKeyBytes() (privateBytes, publicBytes []byte, err error) {
	pri, pub, err := generateKey()
	if err != nil {
		return nil, nil, err
	}
	priBytes, err := x509.MarshalPKCS8PrivateKey(pri)
	if err != nil {
		return nil, nil, err
	}
	pubBytes := x509.MarshalPKCS1PublicKey(pub)
	return priBytes, pubBytes, nil
}

func GenerateKey64() (pri64, pub64 string, err error) {
	pri, pub, err := GenerateKeyBytes()
	if err != nil {
		return "", "", nil
	}
	return base64.StdEncoding.EncodeToString(pri),
		base64.StdEncoding.EncodeToString(pub),
		nil
}

func PublicKeyFrom(key []byte) (*rsa.PublicKey, error) {
	pub, err := x509.ParsePKCS1PublicKey(key)
	if err != nil {
		return nil, err
	}
	return pub, nil
}

func PublicKeyFrom64(key string) (*rsa.PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return PublicKeyFrom(b)
}

func PrivateKeyFrom(key []byte) (*rsa.PrivateKey, error) {
	pri, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		return nil, err
	}
	p, ok := pri.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("invalid private key")
	}
	return p, nil
}

func PrivateKeyFrom64(key string) (*rsa.PrivateKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFrom(b)
}

func PublicEncrypt(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, key, data)
}

func PublicEncrypt64(pubkey64 string, data string) (string, error) {
	publicKey, err := PublicKeyFrom64(pubkey64)
	if err != nil {
		return "", err
	}

	//encrypt data
	cryptBytes, err := PublicEncrypt(publicKey, []byte(data))
	if err != nil {
		return "", err
	}
	//convert to base64 and return
	return base64.StdEncoding.EncodeToString(cryptBytes), nil
}

func PublicSign(key *rsa.PublicKey, data []byte) ([]byte, error) {
	return PublicEncrypt(key, hash(data))
}

func PublicVerify(key *rsa.PublicKey, sign, data []byte) error {
	return rsa.VerifyPKCS1v15(key, crypto.SHA1, hash(data), sign)
}

func PrivateDecrypt(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, key, data)
}

func PrivateDecrypt64(prikey64 string, data string) (string, error) {
	privateKey, err := PrivateKeyFrom64(prikey64)
	if err != nil {
		return "", err
	}

	dataBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	//decrypt data
	decryptBytes, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, dataBytes)
	if err != nil {
		return "", err
	}
	return string(decryptBytes), nil
}

func PrivateSign(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA1, hash(data))
}

func PrivateVerify(key *rsa.PrivateKey, sign, data []byte) error {
	h, err := PrivateDecrypt(key, sign)
	if err != nil {
		return err
	}
	if !bytes.Equal(h, hash(data)) {
		return rsa.ErrVerification
	}
	return nil
}
