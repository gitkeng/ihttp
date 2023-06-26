package rsautil_test

import (
	"testing"

	"github.com/gitkeng/ihttp/util/cryptutil/rsautil"
	"github.com/magiconair/properties/assert"
)

func TestSSHKeyPairUtil(t *testing.T) {
	privateKeyBytes, publicKeyBytes, err := rsautil.GenerateSSHKeyPair()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Private Key :\n%s\n", string(privateKeyBytes))
	t.Logf("Public Key :\n%s\n", string(publicKeyBytes))

}

func TestEncryptDecrypt(t *testing.T) {
	privateKey64, publicKey64, err := rsautil.GenerateKey64()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Private Key :\n%s\n", privateKey64)
	t.Logf("Public Key :\n%s\n", publicKey64)

	plainText := "test data"

	//encrypt with public key
	cryptData, err := rsautil.PublicEncrypt64(publicKey64, plainText)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Crypt data :\n%s\n", cryptData)

	//decrypt with private key
	decryptResult, err := rsautil.PrivateDecrypt64(privateKey64, cryptData)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("decrypt data :\n%s\n", decryptResult)
	assert.Equal(t, plainText, decryptResult)
}
