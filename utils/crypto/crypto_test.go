package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateKeys(t *testing.T) {
	privateKey, publicKey, _ := GenerateKeys()
	assert.Equal(t, 32, cap(privateKey))
	assert.Equal(t, 33, cap(publicKey))
}

func TestCombinePublicKeys(t *testing.T) {
	keys := [][]byte{TestPublicKey1, TestPublicKey2, TestPublicKey3}
	aggregate, err := CombinePublicKeys(keys)
	assert.Nil(t, err)
	assert.Equal(t, 33, cap(aggregate))
}

func TestCombineSignatures(t *testing.T) {
	combinedSignature, err := CombineSignatures([][]byte{TestSignature1,
		TestSignature2, TestSignature3})
	assert.Nil(t, err)
	assert.NotNil(t, combinedSignature)
}

func TestVerifySingleSignature(t *testing.T) {
	signature, err := SignMessage(TestMessage, TestPrivateKey1)
	assert.Nil(t, err)
	verification, err := VerifySignature(TestMessage, signature, TestPublicKey1)
	assert.Nil(t, err)
	assert.True(t, verification)
}

func TestVerifyMultiplePartialSignature(t *testing.T) {
	combinedSignature, err := CombineSignatures([][]byte{TestSignature1,
		TestSignature2, TestSignature3})
	assert.Nil(t, err)
	assert.NotNil(t, combinedSignature)
	verification, err := VerifySignature(TestMessage, combinedSignature, CombinedPublicKey)
	assert.Nil(t, err)
	assert.True(t, verification)
}
